package smsproc

import (
	"context"
	"encoding/json"
	"errors"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"sms/modelq"
	"sms/smsru"
	"time"
)

const (
	smsType = "sms"
)

var ErrUnknownType = errors.New("unknown notification type")

type Processor struct {
	amqpURL   string
	queueName string
	cli       *smsru.Client
}

func NewSmsProcessor(amqpURL string, queueName string, cli *smsru.Client) *Processor {
	return &Processor{
		amqpURL:   amqpURL,
		queueName: queueName,
		cli:       cli,
	}
}

func (p *Processor) Start(ctx context.Context) {
connloop:
	for {
		select {
		case <-ctx.Done():
			break connloop
		default:
			// trying to connect
		}

		err := p.connectConsume(ctx)
		if err != nil {
			log.WithField("err", err).Error("consuming exited with error")
			time.Sleep(2 * time.Second)
		}
	}
}

func (p *Processor) connectConsume(ctx context.Context) error {
	log.Info("Connecting to amqp")
	conn, err := amqp.Dial(p.amqpURL)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	notifyCh, err := p.consume(ch, p.queueName)
	if err != nil {
		return err
	}

	finishedCh := make(chan struct{}, 1)

	go func() {
		log.WithField("queueName", p.queueName).Info("Accepting messages")

		for msg := range notifyCh {
			var message modelq.Notification

			err := json.Unmarshal(msg.Body, &message)
			if err != nil {
				log.WithField("err", err).Error("response unmarshal fail")
				msg.Nack(true, true)
				continue
			}

			err = p.process(message)
			if err != nil {
				log.WithField("err", err).Error("process fail")
				msg.Nack(true, true)
				continue
			}

			msg.Ack(false)
		}

		finishedCh <- struct{}{}
	}()

	select {
	case <-finishedCh:
		return nil
	case <-ctx.Done():
		return nil
	}
}

func (p *Processor) consume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
	queue, err := ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		return nil, err
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		false,      // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return nil, err
	}

	return msgs, nil
}

func (p *Processor) process(msg modelq.Notification) error {
	if msg.Type != smsType {
		return ErrUnknownType
	}

	resp, err := p.cli.Send(msg.Recipient, msg.Content)
	log.WithError(err).WithField("resp", resp).Info("sent sms")

	return err
}