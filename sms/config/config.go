package config

import "github.com/caarlos0/env"

type Config struct {
	SmsHost   string `env:"SMS_HOST" envDefault:"https://sms.ru"`
	SmsApiID  string `env:"SMS_API_ID"`
	AmqpURL   string `env:"AMQP_URL" envDefault:"amqp://guest:guest@rabbit:5672/"`
	QueueName string `env:"QUEUE_NAME" envDefault:"notifications"`
}

func EnvConfig() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
