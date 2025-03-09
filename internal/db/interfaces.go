package db

import (
	"context"
	"github.com/MrPixik/url_shortener/internal/app/models"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock_db_service.go -package=mocks DatabaseService
type DatabaseService interface {
	Ping() error
	CreateUrl(ctx context.Context, shortURL, originalURL string, userId int) error
	CreateUrls(ctx context.Context, urls []models.URLMapping, userId int) error
	GetUrlByShortName(ctx context.Context, shortUrl string, userId int) (models.UrlsObj, error)
	GetUrlsByUserId(ctx context.Context, userId int) ([]models.URLMapping, error)
	CreateUser(ctx context.Context, login, password string) error
	AuthenticateUser(ctx context.Context, login, password string) (int, error)
}
