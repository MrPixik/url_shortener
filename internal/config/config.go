package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/spf13/pflag"
)

// Config struct for server configuration
type Config struct {
	LocalServerAddr string `env:"SERVER_ADDRESS"`
	ShortURLAddr    string `env:"BASE_URL"`
}

// InitConfig func for config initialization
func InitConfig() *Config {
	config := &Config{}

	env.Parse(config)

	if config.LocalServerAddr == "" {
		pflag.StringVar(&config.LocalServerAddr, "a", "localhost:8080", "local server address")
	}

	if config.ShortURLAddr == "" {
		pflag.StringVar(&config.ShortURLAddr, "b", "localhost:8080", "Short URL address")
	}

	pflag.Parse()

	return config
}
