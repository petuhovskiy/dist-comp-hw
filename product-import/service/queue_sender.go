package service

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"product-import/modelq"
	"product-import/qsession"
)

type QueueSender struct {
	q *qsession.Session
}

func NewQueueSender(amqpURL string, queue string) *QueueSender {
	q := qsession.New(queue, amqpURL, false)

	return &QueueSender{
		q: q,
	}
}

func (n *QueueSender) Send(t modelq.ProductImport) error {
	body, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return n.q.Push(amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}
