package config

import (
	"fmt"
	"os"
	"strings"
)

type (
	Config struct {
		HTTPPort      string
		DatabaseURL   string
		JWTSecret     string
		AdminEmail    string
		AdminPassword string
	}
)

func Load() (Config, error) {
	var missing []string

	cfg := Config{
		HTTPPort:      env("HTTP_PORT", &missing),
		DatabaseURL:   env("DATABASE_URL", &missing),
		JWTSecret:     env("JWT_SECRET", &missing),
		AdminEmail:    env("ADMIN_EMAIL", &missing),
		AdminPassword: env("ADMIN_PASSWORD", &missing),
	}

	if len(missing) > 0 {
		return Config{}, fmt.Errorf("missing required environment variables: %s", strings.Join(missing, ", "))
	}

	return cfg, nil
}

func env(key string, missing *[]string) string {
	value := os.Getenv(key)
	if value == "" {
		*missing = append(*missing, key)
	}
	return value
}
