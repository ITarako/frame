package config

import (
	"github.com/caarlos0/env/v9"
)

type DBConfig struct {
	Host         string `env:"POSTGRES_HOST"`
	Port         string `env:"POSTGRES_PORT"`
	User         string `env:"POSTGRES_USER"`
	Password     string `env:"POSTGRES_PASSWORD"`
	Name         string `env:"POSTGRES_DB"`
	MaxOpenConns int    `env:"POSTGRES_MAX_OPEN_CONNS"`
	MaxIdleConns int    `env:"POSTGRES_MAX_IDLE_CONNS"`
	MaxIdleTime  string `env:"POSTGRES_MAX_IDLE_TIME"`
	SslMode      string `env:"POSTGRES_SSL_MODE"`
}

func NewDBConfig() (*DBConfig, error) {
	cfg := &DBConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
