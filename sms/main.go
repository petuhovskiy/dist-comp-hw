package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"sms/config"
	"sms/smsproc"
	"sms/smsru"
)

func main() {
	conf, err := config.EnvConfig()
	if err != nil {
		log.WithError(err).Fatal("failed to read config")
	}

	smsClient := smsru.NewClient(conf.SmsHost, conf.SmsApiID)

	proc := smsproc.NewSmsProcessor(conf.AmqpURL, conf.QueueName, smsClient)
	proc.Start(context.Background())
}
