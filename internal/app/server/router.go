package server

import (
	"github.com/MrPixik/url_shortener/internal/app/middleware"
	"github.com/MrPixik/url_shortener/internal/config"
	"github.com/MrPixik/url_shortener/internal/db"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

// InitHandlers func for creating new chi.Router with all Handlers
func InitHandlers(cfg *config.Config, logger *zap.SugaredLogger, db db.DatabaseService) chi.Router {
	router := chi.NewRouter()

	//Global middleware
	router.Use(
		middleware.LoggingMiddleware(logger),
		middleware.CompressingMiddleware)

	//Public endpoints
	router.Post("/register", wrap(registrationPostHandler, cfg, db))
	router.Post("/login", wrap(loginPostHandler, cfg, db))
	router.Get("/ping", wrap(pingDBHandler, cfg, db))

	//Protected endpoints
	router.Route("/", func(router chi.Router) {
		router.Use(middleware.AuthenticationMiddleware)

		router.Get("/{shortURL}", wrap(mainPageGetHandler, cfg, db))
		router.Post("/", wrap(mainPagePostHandler, cfg, db))

		router.Route("/api", func(api chi.Router) {
			api.Post("/shorten", wrap(shortenURLPostHandler, cfg, db))
			api.Post("/shorten/batch", wrap(urlBatchPostHandler, cfg, db))
			api.Get("/user/urls", wrap(userGetHandler, cfg, db))
		})
	})

	return router
}

func wrap(handler func(w http.ResponseWriter, r *http.Request, cfg *config.Config, db db.DatabaseService), cfg *config.Config, db db.DatabaseService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(w, r, cfg, db)
	}
}
