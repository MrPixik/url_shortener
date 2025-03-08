package db

import (
	"context"
	"github.com/MrPixik/url_shortener/internal/app/models"
)

//go:generate mockgen -source=interfaces.go -destination=mocks/mock_db_service.go -package=mocks DatabaseService
type DatabaseService interface {
	Ping() error
	CreateUrl(ctx context.Context, shortURL, originalURL string) error
	CreateUrls(ctx context.Context, urls []models.URLMapping) error
	GetUrlByShortName(ctx context.Context, shortUrl string) (models.URLDB, error)
	CreateUser(ctx context.Context, login, password string) error
	AuthenticateUser(ctx context.Context, login, password string) (bool, error)
}
