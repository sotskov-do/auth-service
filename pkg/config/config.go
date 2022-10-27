package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Server serverConfig
}

type serverConfig struct {
	Port      string `envconfig:"PORT" default:"8080"`
	PgUrl     string `envconfig:"PG_URL" default:"postgres://postgres:1111@localhost:5432/postgres"`
	SecretKey string `envconfig:"SECRET_KEY" default:"secret key"`
}

func New() (*Config, error) {
	var c Config

	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
