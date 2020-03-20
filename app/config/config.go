package config

import "github.com/caarlos0/env"

type Config struct {
	BindAddr     string `env:"BIND_ADDR" envDefault:":8080"`
	PostgresAddr string `env:"POSTGRES_ADDR" envDefault:"host=localhost port=5432 user=postgres password=DB_PASSWORD dbname=postgres sslmode=disable"` //nolint
}

func EnvConfig() (*Config, error) {
	var config Config
	err := env.Parse(&config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
