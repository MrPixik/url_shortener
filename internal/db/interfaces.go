package db

import "github.com/MrPixik/url_shortener/internal/app/models/easyjson"

type DatabaseService interface {
	Ping() error
	CreateUrl(shortURL, originalURL string) error
	GetUrlByShortName(shortUrl string) (easyjson.URLDB, error)
}
