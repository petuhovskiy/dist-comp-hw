package batchimport

import (
	"app/modeldb"
	"app/modelq"
	"app/repos/psql"
	"context"
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"time"
)

type ImportWatcher struct {
	amqpURL   string
	queueName string
	repo      *psql.Products
}

func NewImportWatcher(amqpURL string, queueName string, repo *psql.Products) *ImportWatcher {
	return &ImportWatcher{
		amqpURL:   amqpURL,
		queueName: queueName,
		repo:      repo,
	}
}

func (p *ImportWatcher) Start(ctx context.Context) {
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

func (p *ImportWatcher) connectConsume(ctx context.Context) error {
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
			var message modelq.ProductImport

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

func (p *ImportWatcher) consume(ch *amqp.Channel, queueName string) (<-chan amqp.Delivery, error) {
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

func (p *ImportWatcher) process(msg modelq.ProductImport) error {
	var products []modeldb.Product
	for _, prod := range msg.Products {
		products = append(products, modeldb.Product{
			ID:       0,
			Name:     prod.Name,
			Code:     prod.Code,
			Category: prod.Category,
		})
	}

	rows, err := p.repo.InsertMany(products)
	log.WithError(err).WithField("rows", rows).WithField("products", len(products)).Info("inserted products")

	return err
}
