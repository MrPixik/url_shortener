package db

import (
	"database/sql"
)

type Storage struct {
	db *sql.DB
}

func (storage *Storage) Ping() error {
	return storage.db.Ping()
}
