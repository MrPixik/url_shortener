package config

import (
	"github.com/caarlos0/env/v11"
	"github.com/spf13/pflag"
)

// Config struct for server configuration
type Config struct {
	LocalServerAddr string `env:"SERVER_ADDRESS"`
	ShortURLAddr    string `env:"BASE_URL"`
	FileStoragePath string `env:"FILE_STORAGE_PATH"`
	DatabaseDSN     string `env:"DATABASE_DSN"`
}

// InitConfig func for config initialization
func InitConfig() (*Config, error) {
	cfg := &Config{}

	err := env.Parse(cfg)
	if err != nil {
		return
	}

	if cfg.LocalServerAddr == "" {
		pflag.StringVar(&cfg.LocalServerAddr, "a", "localhost:8080", "local server address")
	}

	if cfg.ShortURLAddr == "" {
		pflag.StringVar(&cfg.ShortURLAddr, "b", "localhost:8080", "Short OrigURL address")
	}

	if cfg.FileStoragePath == "" {
		pflag.StringVar(&cfg.FileStoragePath, "f", "./tmp/short-url-db.json", "File storage path")
	}
	if cfg.DatabaseDSN == "" {
		pflag.StringVar(&cfg.DatabaseDSN, "d", "user=postgres password=admin host=localhost port=5432 dbname=url_shortener_db sslmode=disable", "Data Source Name")
	}

	pflag.Parse()
}
