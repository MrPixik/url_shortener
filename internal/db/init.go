package db

import (
	"database/sql"
	"github.com/MrPixik/url_shortener/internal/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"go.uber.org/zap"
)

const DRIVERNAME = "pgx"

func InitDBService(cfg *config.Config, logger *zap.SugaredLogger) *Storage {
	dbService, err := connect(cfg)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Infoln("Database connected successfully")
	return dbService
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
