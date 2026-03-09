package configs

import (
	"fmt"
	"time"
)

type Config struct {
	SERVER_PORT string `env:"SERVER_PORT"`
	DB_PASSWORD string `env:"DB_PASSWORD"`
	DB_USERNAME string `env:"DB_USERNAME"`
	DB_HOST     string `env:"DB_HOST"`
	DB_PORT     string `env:"DB_PORT"`
	DB_NAME     string `env:"DB_NAME"`
	InternalConfig
}

type InternalConfig struct {
	SALT            string `env:"SALT"`
	JWT_SIGNING_KEY string 	`env:"JWT_SIGNING_KEY"`
	TOKEN_TTL		time.Duration `env:"TOKEN_TTL"`
}

func (c *Config) DSN() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%s/%s", c.DB_USERNAME, c.DB_PASSWORD, c.DB_HOST, c.DB_PORT, c.DB_NAME)
}
