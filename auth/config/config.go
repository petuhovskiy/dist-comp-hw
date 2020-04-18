package config

import (
	"github.com/caarlos0/env"
	"time"
)

type Config struct {
	BindAddr     string `env:"BIND_ADDR" envDefault:":8081"`
	PostgresAddr string `env:"POSTGRES_ADDR" envDefault:"host=localhost port=5432 user=postgres password=DB_PASSWORD dbname=postgres sslmode=disable"` //nolint

	AccessTokenExpiration  time.Duration `env:"ACCESS_TOKEN_EXPR" envDefault:"3m"`
	RefreshTokenExpiration time.Duration `env:"REFRESH_TOKEN_EXPR" envDefault:"10m"`
	TokenLength            uint          `env:"TOKEN_LENGTH" envDefault:"32"`
	JWTSecret              string        `env:"JWT_SECRET"`

	AmqpURL    string `env:"AMQP_URL" envDefault:"amqp://guest:guest@rabbit:5672/"`
	QueueSms   string `env:"QUEUE_SMS" envDefault:"notifications.sms"`
	QueueEmail string `env:"QUEUE_EMAIL" envDefault:"notifications.email"`
	ConfirmUrl string `env:"CONFIRM_URL" envDefault:"http://localhost:8081/v1/confirm?v=%s"`

	Grpc Grpc
}

type Grpc struct {
	Bind    string `env:"GRPC_BIND" envDefault:":9090"`
	Timeout int    `env:"GRPC_TIMEOUT" envDefault:"5"`
}

func EnvConfig() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
