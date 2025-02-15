package db

import (
	"database/sql"
	"github.com/MrPixik/url_shortener/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
)

type Storage struct {
	db *sql.DB
}

func connect(cfg *config.Config) (*Storage, error) {

	db, err := sql.Open(DRIVERNAME, cfg.DatabaseDSN)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil
}

func (storage *Storage) Ping() error {
	return storage.db.Ping()
}
