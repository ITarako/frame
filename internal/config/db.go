package config

import (
	"github.com/caarlos0/env/v9"
)

type DBConfig struct {
	Host         string `env:"DB_HOST"`
	Port         string `env:"DB_PORT"`
	User         string `env:"DB_USER"`
	Password     string `env:"DB_PASSWORD"`
	Name         string `env:"DB_NAME"`
	MaxOpenConns int    `env:"DB_MAX_OPEN_CONNS"`
	MaxIdleConns int    `env:"DB_MAX_IDLE_CONNS"`
	MaxIdleTime  string `env:"DB_MAX_IDLE_TIME"`
	SslMode      string `env:"DB_SSL_MODE"`
}

func NewDBConfig() (*DBConfig, error) {
	cfg := &DBConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
