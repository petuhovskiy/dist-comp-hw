package main

import (
	"context"
	"email/config"
	"email/emailproc"
	"email/sendmail"
	log "github.com/sirupsen/logrus"
)

func main() {
	conf, err := config.EnvConfig()
	if err != nil {
		log.WithError(err).Fatal("failed to read config")
	}

	sender := sendmail.NewSender(conf.Email)

	proc := emailproc.NewEmailProcessor(conf.AmqpURL, conf.QueueName, sender)
	proc.Start(context.Background())
}
