package config

import (
	"errors"
	"flag"
	"github.com/joho/godotenv"
)

func Load() error {
	var path string
	flag.StringVar(&path, "env-file", ".env", "Env file (must be place in root of project)")
	flag.Parse()

	if path == "" {
		return errors.New("env file is not set")
	}

	err := godotenv.Load(path)
	if err != nil {
		return err
	}

	return nil
}
