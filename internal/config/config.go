package config

import (
	"errors"
	"log/slog"

	"github.com/caarlos0/env"
)

var errServerEmpty = errors.New("server empty")

// Config конфигурация сервера
type Config struct {
	Level slog.Level `env:"LOG_LEVEL"`
	Server string `env:"SERVER_ADDRESS"`
	DataBaseDSN string `env:"DATABASE_DSN"`
}

// GetConfig получить конфигурацию сервера
func (c *Config) GetConfig() error {
	if err := env.Parse(c); err != nil {
		return err
	}
	if c.Server == "" {
		return errServerEmpty
	}
	return nil
}