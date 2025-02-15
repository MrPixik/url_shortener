package db

import "github.com/MrPixik/url_shortener/internal/config"

const DRIVERNAME = "pgx"

func InitDBService(cfg *config.Config) (*Storage, error) {
	return connect(cfg)
}
