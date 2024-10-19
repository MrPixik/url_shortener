package server

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/MrPixik/url_shortener/internal/app/middleware"
	"github.com/MrPixik/url_shortener/internal/app/models"
	easyjson2 "github.com/MrPixik/url_shortener/internal/app/models/easyjson"
	"github.com/MrPixik/url_shortener/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"io"
	"net/http"
	"os"
)

func generateShortUrl(longUrl string) string {
	hasher := md5.New()
	shortURL := hex.EncodeToString(hasher.Sum([]byte(longUrl))[0:12])
	return shortURL
}

// InitHandlers func for creating new chi.Router with all Handlers
func InitHandlers(cfg *config.Config, logger *zap.SugaredLogger) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.LoggingMiddleware(logger), middleware.CompressingMiddleware)

	router.Route("/", func(r chi.Router) {
		r.Get("/{id}", func(w http.ResponseWriter, r *http.Request) {
			mainPageGetHandler(w, r, cfg)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			mainPagePostHandler(w, r, cfg)
		})
		r.Post("/api/shorten", func(w http.ResponseWriter, r *http.Request) {
			shortenURLPostHandler(w, r, cfg)
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
	shortURL := generateShortUrl(URL)
	//fmt.Println("\"" + "http://localhost:8080/" + shortURL + "\"" + " post")

	//Writing to file
	fileHandler, err := models.NewFileHandler(cfg.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = fileHandler.WriteURLToFile(&easyjson2.URLFileRecord{
		Original: URL,
		Short:    shortURL,
	})
	if err != nil {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://" + cfg.ShortURLAddr + "/" + shortURL))
}

func shortenURLPostHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {

	var urlReq easyjson2.URLRequest
	err := easyjson.UnmarshalFromReader(r.Body, &urlReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fmt.Println("Received URL:", URL)
	shortURL := generateShortUrl(urlReq.URL)

	urlRes := easyjson2.URLResponse{
		URL: "http://" + cfg.ShortURLAddr + "/" + shortURL,
	}

	//Writing to file
	fileHandler, err := models.NewFileHandler(cfg.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	err = fileHandler.WriteURLToFile(&easyjson2.URLFileRecord{
		Original: urlReq.URL,
		Short:    shortURL,
	})
	if err != nil {
		return
	}
	//fmt.Println("\"" + "http://localhost:8080/" + shortURL + "\"" + " post")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err = easyjson.MarshalToWriter(urlRes, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func mainPageGetHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {
	//fmt.Println("\"" + chi.URLParam(r, "id") + "\"" + " get")
	reqShortURL := chi.URLParam(r, "id")
	fileHandler, err := models.NewFileHandler(cfg.FileStoragePath, os.O_RDONLY|os.O_CREATE)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	for {
		record, err := fileHandler.ReadURLFromFile()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if record.Short == reqShortURL {
			w.Header().Set("Location", record.Original)
			w.WriteHeader(http.StatusTemporaryRedirect)
			return
		}

	}
	w.WriteHeader(http.StatusBadRequest)

}
