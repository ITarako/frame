package config

import (
	"flag"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

const (
	EnvLocal = "local"
	EnvTest  = "test"
	EnvProd  = "prod"
)

type Project struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
}

type Server struct {
	Host         string `yaml:"host"`
	Port         int    `yaml:"port"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
	IdleTimeout  int    `yaml:"idle_timeout"`
}

type Database struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	Name         string `yaml:"name"`
	MaxOpenConns int    `yaml:"max_open_conns"`
	MaxIdleConns int    `yaml:"max_idle_conns"`
	MaxIdleTime  string `yaml:"max_idle_time"`
	SslMode      string `yaml:"sslmode"`
}

type Config struct {
	Project  Project  `yaml:"project"`
	Server   Server   `yaml:"server"`
	Database Database `yaml:"database"`
}

func Parse() *Config {
	var configPath string
	flag.StringVar(&configPath, "config-path", "config.yml", "Config path")
	flag.Parse()

	if configPath == "" {
		log.Fatal("CONFIG_PATH is not set")
	}

	configPath = filepath.Clean(configPath)
	file, err := os.Open(configPath)
	if err != nil {
		log.Fatalf("cannot read file: %s, err: %s\n", configPath, err)
	}
	defer func() {
		_ = file.Close()
	}()

	cfg := &Config{}
	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(cfg); err != nil {
		log.Fatalf("cannot read config: %s\n", err)
	}

	return cfg
}
