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
	BaseURL     string
	DatabaseURL string
	Facebook    FacebookConfig
}

// FacebookConfig holds Facebook OAuth settings.
type FacebookConfig struct {
	ClientID     string
	ClientSecret string
}

// Load reads configuration from environment variables and flags.
// Secrets come from environment variables; non-secret config from flags.
func Load() (Config, error) {
	var cfg Config

	flag.StringVar(&cfg.Addr, "addr", ":8080", "HTTP listen address")
	flag.StringVar(&cfg.BaseURL, "base-url", "http://localhost:8080", "Base URL for OAuth callbacks")
	flag.Parse()

	cfg.DatabaseURL = os.Getenv("DATABASE_URL")
	if cfg.DatabaseURL == "" {
		return Config{}, fmt.Errorf("DATABASE_URL environment variable is required")
	}

	cfg.Facebook.ClientID = os.Getenv("FACEBOOK_CLIENT_ID")
	cfg.Facebook.ClientSecret = os.Getenv("FACEBOOK_CLIENT_SECRET")

	return cfg, nil
}

// HasFacebookAuth returns true if Facebook OAuth is configured.
func (c *Config) HasFacebookAuth() bool {
	return c.Facebook.ClientID != "" && c.Facebook.ClientSecret != ""
}
