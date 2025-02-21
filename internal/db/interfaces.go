package db

import "github.com/MrPixik/url_shortener/internal/app/models/easyjson"

//go:generate mockgen -source=interfaces.go -destination=mocks/mock_db_service.go -package=mocks DatabaseService
type DatabaseService interface {
	Ping() error
	CreateUrl(shortURL, originalURL string) error
	GetUrlByShortName(shortUrl string) (easyjson.URLDB, error)
}
