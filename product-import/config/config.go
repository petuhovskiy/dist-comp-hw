package config

import (
	"github.com/caarlos0/env"
)

type Config struct {
	BindAddr string `env:"BIND_ADDR" envDefault:":8082"`

	AmqpURL     string `env:"AMQP_URL" envDefault:"amqp://guest:guest@rabbit:5672/"`
	QueueImport string `env:"QUEUE_IMPORT" envDefault:"products.import"`
}

func EnvConfig() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
