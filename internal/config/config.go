// Package config handles application configuration from environment
// variables and command-line flags.
package config

import (
	"flag"
	"fmt"
	"os"
)

// Config holds all application configuration.
type Config struct {
	Addr        string
	DatabaseURL string
}

// Load reads configuration from environment variables and flags.
// Secrets come from environment variables; non-secret config from flags.
func Load() (Config, error) {
	var cfg Config

	flag.StringVar(&cfg.Addr, "addr", ":8080", "HTTP listen address")
	flag.Parse()

	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL environment variable is required")
	}

	return cfg, nil
}
