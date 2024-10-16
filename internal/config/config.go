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

var Cfg *Config

// InitConfig func for config initialization
func InitConfig() {
	cfg := &Config{}
	Cfg = cfg

	env.Parse(Cfg)

	if Cfg.LocalServerAddr == "" {
		pflag.StringVar(&Cfg.LocalServerAddr, "a", "localhost:8080", "local server address")
	}

	if Cfg.ShortURLAddr == "" {
		pflag.StringVar(&Cfg.ShortURLAddr, "b", "localhost:8080", "Short URL address")
	}

	pflag.Parse()
}
