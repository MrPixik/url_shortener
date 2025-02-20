package db

import (
	"database/sql"
	"github.com/MrPixik/url_shortener/internal/app/models/easyjson"
	"io"
	"os"
)

const INITQUERYPATH = "internal/db/skripts/initTables.sql"

type Storage struct {
	db *sql.DB
}

func (storage *Storage) Ping() error {
	return storage.db.Ping()
}

func (storage *Storage) init() error {
	initQuery, err := readStringFromFile(INITQUERYPATH)
	if err != nil {
		return err
	}
	if _, err := storage.db.Exec(initQuery); err != nil {
		return err
	}
	return nil
}

func (storage *Storage) CreateUrl(shortURL, originalURL string) error {
	query := `INSERT INTO url (short_url, long_url) VALUES ($1,$2)`

	if _, err := storage.db.Exec(query, shortURL, originalURL); err != nil {
		return err
	}
	return nil
}
func (storage *Storage) GetUrlByShortName(shortUrl string) (easyjson.URLDB, error) {
	query := `SELECT url_id, long_url FROM url WHERE short_url = $1`

	url := easyjson.URLDB{Short: shortUrl}

	err := storage.db.QueryRow(query, shortUrl).Scan(&url.ID, &url.Original)

	return url, err
}

func readStringFromFile(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 2228)
	if err != nil {
		return "", err
	}

	body, err := io.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(body), nil
}
