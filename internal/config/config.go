package config

import "github.com/spf13/pflag"

// Config struct for server configuration
type Config struct {
	LocalServerAddr string
	ShortURLAddr    string
}

// InitConfig func for config initialization
func InitConfig() *Config {
	config := &Config{}

	pflag.StringVar(&config.LocalServerAddr, "a", "localhost:8080", "local server address")
	pflag.StringVar(&config.ShortURLAddr, "b", "localhost:8080", "Short URL address")

	pflag.Parse()

	return config
}
