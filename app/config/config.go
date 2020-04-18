package config

import "github.com/caarlos0/env"

type Config struct {
	BindAddr     string `env:"BIND_ADDR" envDefault:":8080"`
	PostgresAddr string `env:"POSTGRES_ADDR" envDefault:"host=localhost port=5432 user=postgres password=DB_PASSWORD dbname=postgres sslmode=disable"` //nolint
	AuthGrpc     string `env:"AUTH_GRPC" envDefault:"auth:9090"`

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
