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
}

func EnvConfig() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
