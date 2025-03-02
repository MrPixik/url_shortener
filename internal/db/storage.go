package db

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/MrPixik/url_shortener/internal/app/models"
	"io"
	"os"
	"strings"
	"time"
)

const (
	INITQUERYPATH  = "internal/db/skripts/initTables.sql"
	REQUESTMAXTIME = time.Second * 3
)

// Storage type implements interaction with database
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

func (storage *Storage) CreateUrl(ctx context.Context, shortURL, originalURL string) error {
	query := `INSERT INTO url (short_url, long_url) VALUES ($1,$2)`

	ctx, cancel := context.WithTimeout(ctx, REQUESTMAXTIME)
	defer cancel()
	if _, err := storage.db.ExecContext(ctx, query, shortURL, originalURL); err != nil {
		return err
	}
	return nil
}

func (storage *Storage) CreateUrls(ctx context.Context, urls []models.URLMapping) error {
	if len(urls) == 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, REQUESTMAXTIME)
	defer cancel()

	// Initialization of placeholders like ($1,$2)...
	placeholders := make([]string, len(urls))
	values := make([]interface{}, 0, len(urls)*2)

	for i, v := range urls {
		placeholders[i] = fmt.Sprintf("($%d,$%d)", 2*i+1, 2*(i+1))
		values = append(values, v.ShortURL, v.OrigURL)
	}

	query := fmt.Sprintf(`INSERT INTO url (short_url, long_url) VALUES %s`, strings.Join(placeholders, ","))

	if _, err := storage.db.ExecContext(ctx, query, values...); err != nil {
		return err
	}
	return nil
}

func (storage *Storage) GetUrlByShortName(ctx context.Context, shortUrl string) (models.URLDB, error) {
	query := `SELECT url_id, short_url, long_url FROM url WHERE short_url = $1`

	url := models.URLDB{}

	ctx, cancel := context.WithTimeout(ctx, REQUESTMAXTIME)
	defer cancel()
	err := storage.db.QueryRowContext(ctx, query, shortUrl).Scan(&url.ID, &url.Short, &url.Original)

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
