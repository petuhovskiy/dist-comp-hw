package qsession

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
	"time"

	"github.com/streadway/amqp"
)

// This exports a Session object that wraps this library. It
// automatically reconnects when the connection fails, and
// blocks all pushes until the connection succeeds. It also
// confirms every outgoing message, so none are lost.
// It doesn't automatically ack each message, but leaves that
// to the parent process, since it is usage-dependent.
//
// Try running this in one terminal, and `rabbitmq-server` in another.
// Stop & restart RabbitMQ to see how the queue reacts.
func Example() {
	name := "job_queue"
	addr := "amqp://guest:guest@localhost:5672/"
	queue := New(name, addr, true)
	message := []byte("message")
	// Attempt to push a message every 3 seconds
	for {
		time.Sleep(time.Second * 3)
		if err := queue.Push(amqp.Publishing{
			ContentType: "text/plain",
			Body:        message,
		}); err != nil {
			fmt.Printf("Push failed: %s\n", err)
		} else {
			fmt.Println("Push succeeded!")
		}
	}
}

// Session provides reliable way to Push messages to RabbitMQ,
// including auto-reconnect and ack confirmations.
type Session struct {
	queueName       string
	logger          log.FieldLogger
	connection      *amqp.Connection
	channel         *amqp.Channel
	done            chan bool
	notifyConnClose chan *amqp.Error
	notifyChanClose chan *amqp.Error
	notifyConfirm   chan amqp.Confirmation
	isReady         bool
	confirmRequired bool
}

const (
	// When reconnecting to the server after connection failure
	reconnectDelay = 5 * time.Second

	// When setting up the channel after a channel exception
	reInitDelay = 2 * time.Second

	// When resending messages the server didn't confirm
	resendDelay = 5 * time.Second
)

var (
	errNotConnected  = errors.New("not connected to a server")
	errAlreadyClosed = errors.New("already closed: not connected to the server")
	errShutdown      = errors.New("session is shutting down")
)

// New creates a new consumer state instance, and automatically
// attempts to connect to the server.
func New(queueName string, addr string, confirmRequired bool) *Session {
	session := Session{
		logger:          log.StandardLogger().WithField("queueName", queueName),
		queueName:       queueName,
		done:            make(chan bool),
		confirmRequired: confirmRequired,
	}
	go session.handleReconnect(addr)
	return &session
}

// handleReconnect will wait for a connection error on
// notifyConnClose, and then continuously attempt to reconnect.
func (session *Session) handleReconnect(addr string) {
	for {
		session.isReady = false
		session.logger.Debug("Attempting to connect")

		conn, err := session.connect(addr)

		if err != nil {
			session.logger.WithField("err", err).Error("Failed to connect. Retrying...")

			select {
			case <-session.done:
				return
			case <-time.After(reconnectDelay):
			}
			continue
		}

		if done := session.handleReInit(conn); done {
			break
		}
	}
}

// connect will create a new AMQP connection
func (session *Session) connect(addr string) (*amqp.Connection, error) {
	conn, err := amqp.Dial(addr)

	if err != nil {
		return nil, err
	}

	session.changeConnection(conn)
	session.logger.Info("Connected!")
	return conn, nil
}

// handleReconnect will wait for a channel error
// and then continuously attempt to re-initialize both channels
func (session *Session) handleReInit(conn *amqp.Connection) bool {
	for {
		session.isReady = false

		err := session.init(conn)

		if err != nil {
			session.logger.WithField("err", err).Error("Failed to initialize channel. Retrying...")

			select {
			case <-session.done:
				return true
			case <-time.After(reInitDelay):
			}
			continue
		}

		select {
		case <-session.done:
			return true
		case <-session.notifyConnClose:
			session.logger.Warn("Connection closed. Reconnecting...")
			return false
		case <-session.notifyChanClose:
			session.logger.Warn("Channel closed. Re-running init...")
		}
	}
}

// init will initialize channel & declare queue
func (session *Session) init(conn *amqp.Connection) error {
	ch, err := conn.Channel()

	if err != nil {
		return err
	}

	if session.confirmRequired {
		err = ch.Confirm(false) // enable confirms only if needed
		if err != nil {
			return err
		}
	}

	_, err = ch.QueueDeclare(
		session.queueName,
		true,  // Durable
		false, // Delete when unused
		false, // Exclusive
		false, // No-wait
		nil,   // Arguments
	)

	if err != nil {
		return err
	}

	session.changeChannel(ch)
	session.isReady = true
	session.logger.Debug("Setup!")

	return nil
}

// changeConnection takes a new connection to the queue,
// and updates the close listener to reflect this.
func (session *Session) changeConnection(connection *amqp.Connection) {
	session.connection = connection
	session.notifyConnClose = make(chan *amqp.Error)
	session.connection.NotifyClose(session.notifyConnClose)
}

// changeChannel takes a new channel to the queue,
// and updates the channel listeners to reflect this.
func (session *Session) changeChannel(channel *amqp.Channel) {
	session.channel = channel
	session.notifyChanClose = make(chan *amqp.Error)
	session.notifyConfirm = make(chan amqp.Confirmation, 1)
	session.channel.NotifyClose(session.notifyChanClose)
	session.channel.NotifyPublish(session.notifyConfirm)
}

// Push will push data onto the queue, and wait for a confirm.
// If no confirms are received until within the resendTimeout,
// it continuously re-sends messages until a confirm is received.
// This will block until the server sends a confirm. Errors are
// only returned if the push action itself fails, see unsafePush.
func (session *Session) Push(publ amqp.Publishing) error {
	if !session.isReady {
		return errors.New("failed to push push: not connected")
	}
	if !session.confirmRequired {
		return session.unsafePush(publ)
	}

	for {
		err := session.unsafePush(publ)
		if err != nil {
			session.logger.WithField("err", err).Warn("Push failed. Retrying...")
			select {
			case <-session.done:
				return errShutdown
			case <-time.After(resendDelay):
			}
			continue
		}
		select {
		case confirm := <-session.notifyConfirm:
			if confirm.Ack {
				session.logger.Info("Push confirmed!")
				return nil
			}
		case <-time.After(resendDelay):
		}
		session.logger.Warn("Push didn't confirm. Retrying...")
	}
}

// unsafePush will push to the queue without checking for
// confirmation. It returns an error if it fails to connect.
// No guarantees are provided for whether the server will
// receive the message.
func (session *Session) unsafePush(publ amqp.Publishing) error {
	if !session.isReady {
		return errNotConnected
	}
	return session.channel.Publish(
		"",                // Exchange
		session.queueName, // Routing key
		false,             // Mandatory
		false,             // Immediate
		publ,
	)
}

// Stream will continuously put queue items on the channel.
// It is required to call delivery.Ack when it has been
// successfully processed, or delivery.Nack when it fails.
// Ignoring this will cause data to build up on the server.
func (session *Session) Stream() (<-chan amqp.Delivery, error) {
	if !session.isReady {
		return nil, errNotConnected
	}
	return session.channel.Consume(
		session.queueName,
		"",    // Consumer
		false, // Auto-Ack
		false, // Exclusive
		false, // No-local
		false, // No-Wait
		nil,   // Args
	)
}

// Close will cleanly shutdown the channel and connection.
func (session *Session) Close() error {
	if !session.isReady {
		return errAlreadyClosed
	}
	err := session.channel.Close()
	if err != nil {
		return err
	}
	err = session.connection.Close()
	if err != nil {
		return err
	}
	close(session.done)
	session.isReady = false
	return nil
}
