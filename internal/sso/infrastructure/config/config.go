package config

import (
	"os"
	"strings"
)

type Config struct {
	DatabaseConnectionString string
	Secret                   string
	Host                     string
	Port                     string
}

type EnvConfig Config

func NewEnvConfig() *EnvConfig {
	config := new(EnvConfig)
	config.load()

	return config
}

func (c *EnvConfig) load() {
	c.DatabaseConnectionString = c.getValue("DB_DSN", "")
	c.Secret = c.getValue("SECRET", "")
	c.Host = c.getValue("HOST", "localhost")
	c.Port = c.getValue("PORT", "8080")
}

func (c *EnvConfig) getValue(key, defaultValue string) string {
	value, ok := os.LookupEnv(key)
	if !ok || strings.TrimSpace(value) == "" {
		return defaultValue
	}

	return strings.TrimSpace(value)
}
