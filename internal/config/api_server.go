package config

import (
	"github.com/caarlos0/env/v9"
)

type APIServerConfig struct {
	Host         string `env:"API_SERVER_HOST"`
	Port         int    `env:"API_SERVER_PORT"`
	WriteTimeout int    `env:"API_SERVER_WRITE_TIMEOUT"`
	ReadTimeout  int    `env:"API_SERVER_READ_TIMEOUT"`
	IdleTimeout  int    `env:"API_SERVER_IDLE_TIMEOUT"`
}

func NewAPIServerConfig() (*APIServerConfig, error) {
	cfg := &APIServerConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
