package config

import "github.com/caarlos0/env/v6"

type Config struct {
	AmqpURL   string `env:"AMQP_URL" envDefault:"amqp://guest:guest@rabbit:5672/"`
	QueueName string `env:"QUEUE_NAME" envDefault:"notifications"`

	Email Email
}

type Email struct {
	Host     string `env:"SMTP_HOST"`
	Port     int    `env:"SMTP_PORT"`
	Username string `env:"SMTP_USERNAME"`
	Password string `env:"SMTP_PASSWORD"`

	From string `env:"EMAIL_FROM"`
}

func EnvConfig() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
