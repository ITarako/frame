package config

import (
	"errors"
	"github.com/caarlos0/env/v9"
	"slices"
)

const (
	EnvLocal = "local"
	EnvTest  = "test"
	EnvProd  = "prod"
)

type ProjectConfig struct {
	Name string `env:"PROJECT_NAME"`
	Env  string `env:"PROJECT_ENV"`
}

func NewProjectConfig() (*ProjectConfig, error) {
	cfg := &ProjectConfig{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	envs := []string{EnvLocal, EnvTest, EnvProd}
	if !slices.Contains(envs, cfg.Env) {
		return nil, errors.New("wrong project environment")
	}

	return cfg, nil
}
