package service

import (
	"auth/modeldb"
	"encoding/json"
	"github.com/streadway/amqp"
	modelq2 "lib/modelq"
	"lib/qsession"
)

type Notificator struct {
	sms *qsession.Session
	email *qsession.Session
}

func NewNotificator(amqpURL string, smsQueue, emailQueue string) *Notificator {
	sms := qsession.New(smsQueue, amqpURL, false)
	email := qsession.New(emailQueue, amqpURL, false)

	return &Notificator{
		sms:   sms,
		email: email,
	}
}

func (n *Notificator) Notify(t modeldb.ConfirmType, recipient string, content string) error {
	var obj interface{}
	var queue *qsession.Session

	switch t {
	case modeldb.ConfirmSms:
		obj = modelq2.Notification{
			Type:      "sms",
			Recipient: recipient,
			Content:   content,
		}
		queue = n.sms

	case modeldb.ConfirmEmail:
		obj = modelq2.Notification{
			Type:      "email",
			Recipient: recipient,
			Content:   content,
		}
		queue = n.email

	default:
		return ErrUnknownNotifyType
	}

	body, err := json.Marshal(obj)
	if err != nil {
		return err
	}

	return queue.Push(amqp.Publishing{
		ContentType: "application/json",
		Body:        body,
	})
}