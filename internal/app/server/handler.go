package server

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"github.com/MrPixik/url_shortener/internal/app/middleware"
	"github.com/MrPixik/url_shortener/internal/config"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"io"
	"net/http"
)

var uRLPool = make(map[string]string)

// InitHandlers func for creating new chi.Router with all Handlers
func InitHandlers(cfg *config.Config, logger *zap.SugaredLogger) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.LoggingMiddleware(logger))

	router.Route("/", func(r chi.Router) {
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			mainPageGetHandler(w, r, cfg)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			mainPagePostHandler(w, r, cfg)
		})
	})
	return router
}

func mainPagePostHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	URL := string(body)
	if URL == "" {
		http.Error(w, "Empty URL", http.StatusBadRequest)
		return
	}

	//fmt.Println("Received URL:", URL)
	hasher := md5.New()
	shortURL := hex.EncodeToString(hasher.Sum(body)[0:12])
	uRLPool[shortURL] = URL
	//fmt.Println("\"" + "http://localhost:8080/" + shortURL + "\"" + " post")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://" + cfg.ShortURLAddr + "/" + shortURL))
}

func mainPageGetHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	fmt.Println("\"" + chi.URLParam(r, "id") + "\"" + " get")
	fmt.Println(uRLPool)
	originalURL, ok := uRLPool[chi.URLParam(r, "id")]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
