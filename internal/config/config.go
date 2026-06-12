package config

import (
	"fmt"
	"os"
)

type DatabaseConfig struct {
	Host     string
	User     string
	Password string
	Name     string
	Port     string
	SSLMode  string
}

func LoadDatabase() DatabaseConfig {
	return DatabaseConfig{
		Host:     env("DB_HOST", "127.0.0.1"),
		User:     env("DB_USER", "postgres"),
		Password: env("DB_PASSWORD", "minhasenha"),
		Name:     env("DB_NAME", "postgres"),
		Port:     env("DB_PORT", "5433"),
		SSLMode:  env("DB_SSLMODE", "disable"),
	}
}

func (c DatabaseConfig) DSN() string {
	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		c.Host, c.User, c.Password, c.Name, c.Port, c.SSLMode,
	)
}

func env(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
