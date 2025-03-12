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

func (storage *Storage) CreateUrl(ctx context.Context, shortURL, originalURL string, userId int) error {
	query := `INSERT INTO urls (short_url, long_url, user_id) VALUES ($1,$2,$3)`

	ctx, cancel := context.WithTimeout(ctx, REQUESTMAXTIME)
	defer cancel()
	if _, err := storage.db.ExecContext(ctx, query, shortURL, originalURL, userId); err != nil {
		return err
	}
	return nil
}

func (storage *Storage) CreateUrls(ctx context.Context, urls []models.URLMapping, userId int) error {
	if len(urls) == 0 {
		return nil
	}

	ctx, cancel := context.WithTimeout(ctx, REQUESTMAXTIME)
	defer cancel()

	// Initialization of placeholders like ($1,$2)...
	placeholders := make([]string, len(urls))
	values := make([]interface{}, 0, len(urls)*2)

	for i, v := range urls {
		placeholders[i] = fmt.Sprintf("($%d,$%d,$%d)", 3*i+1, 3*(i+1), 3*(i+1))
		values = append(values, v.ShortURL, v.OrigURL, userId)
	}

	query := fmt.Sprintf(`INSERT INTO urls (short_url, long_url) VALUES %s`, strings.Join(placeholders, ","))

	if _, err := storage.db.ExecContext(ctx, query, values...); err != nil {
		return err
	}
	return nil
}

func (storage *Storage) GetUrlByShortName(ctx context.Context, shortUrl string, userId int) (models.UrlsObj, error) {
	query := `SELECT url_id, short_url, long_url FROM urls WHERE short_url = $1 AND user_id = $2`

	url := models.UrlsObj{}

	ctx, cancel := context.WithTimeout(ctx, REQUESTMAXTIME)
	defer cancel()
	err := storage.db.QueryRowContext(ctx, query, shortUrl, userId).Scan(&url.ID, &url.Short, &url.Original)

	return url, err
}

func (storage *Storage) GetUrlsByUserId(ctx context.Context, userId int) ([]models.URLMapping, error) {
	query := `SELECT short_url, long_url FROM urls WHERE user_id = $1`

	ctx, cancel := context.WithTimeout(ctx, REQUESTMAXTIME)
	defer cancel()

	rows, err := storage.db.QueryContext(ctx, query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	urlsSlice := make([]models.URLMapping, 0)
	for rows.Next() {
		urls := models.URLMapping{}

		if err = rows.Scan(&urls.ShortURL, &urls.OrigURL); err != nil {
			return nil, err
		}
		urlsSlice = append(urlsSlice, urls)
	}
	return urlsSlice, nil
}

func (storage *Storage) CreateUser(ctx context.Context, login, password string) error {
	query := `INSERT INTO users (login, password) VALUES ($1,$2)`

	ctx, cancel := context.WithTimeout(ctx, REQUESTMAXTIME)
	defer cancel()

	if _, err := storage.db.ExecContext(ctx, query, login, password); err != nil {
		return err
	}
	return nil
}

func (storage *Storage) AuthenticateUser(ctx context.Context, login, password string) (int, error) {
	query := `SELECT user_id FROM users WHERE login = $1 AND password = $2`

	var userId int

	ctx, cancel := context.WithTimeout(ctx, REQUESTMAXTIME)
	defer cancel()
	err := storage.db.QueryRowContext(ctx, query, login, password).Scan(&userId)

	return userId, err
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
