package server

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"github.com/MrPixik/url_shortener/internal/app/middleware"
	easyjson2 "github.com/MrPixik/url_shortener/internal/app/models/easyjson"
	"github.com/MrPixik/url_shortener/internal/config"
	"github.com/MrPixik/url_shortener/internal/db"
	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"io"
	"net/http"
)

const (
	ErrMsgDBWriteError = "An error occurred while writing to database"
	ErrMsgDuplicateURL = "Short URL already exist"
)

func generateShortUrl(longUrl string) string {
	hasher := md5.New()
	return hex.EncodeToString(hasher.Sum([]byte(longUrl))[0:12])
}

// InitHandlers func for creating new chi.Router with all Handlers
func InitHandlers(cfg *config.Config, logger *zap.SugaredLogger, db db.DatabaseService) chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.LoggingMiddleware(logger), middleware.CompressingMiddleware)

	router.Route("/", func(r chi.Router) {
		r.Get("/{shortURL}", func(w http.ResponseWriter, r *http.Request) {
			mainPageGetHandler(w, r, cfg, db)
		})
		r.Post("/", func(w http.ResponseWriter, r *http.Request) {
			mainPagePostHandler(w, r, cfg, db)
		})
		r.Route("/api", func(r chi.Router) {
			r.Post("/shorten", func(w http.ResponseWriter, r *http.Request) {
				shortenURLPostHandler(w, r, cfg, db)
			})
		})
		r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
			pingDBHandler(w, r, cfg, db)

		})
	})
	return router
}

func mainPagePostHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config, db db.DatabaseService) {

	//Reading original URL from request's body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	originalURL := string(body)
	if originalURL == "" {
		http.Error(w, "Empty originalURL", http.StatusBadRequest)
		return
	}
	//Creating short URL
	shortURL := generateShortUrl(originalURL)

	//Creating new object in database
	if err := db.CreateUrl(shortURL, originalURL); err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) {
			if pgErr.Code == "23505" {
				http.Error(w, ErrMsgDuplicateURL, http.StatusConflict)
				return
			}
		}
		http.Error(w, ErrMsgDBWriteError, http.StatusInternalServerError)
		return
	}

	//Writing to file
	//fileHandler, err := models.NewFileHandler(cfg.FileStoragePath, os.O_WRONLY|os.O_CREATE|os.O_APPEND)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//err = fileHandler.WriteURLToFile(&easyjson2.URLFileRecord{
	//	Original: originalURL,
	//	Short:    shortURL,
	//})
	//if err != nil {
	//	return
	//}

	//Configuring response's parameters
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://" + cfg.ShortURLAddr + "/" + shortURL))
}

func shortenURLPostHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config, db db.DatabaseService) {

	//Unmarshalling json from request
	var urlReq easyjson2.URLRequest
	if err := easyjson.UnmarshalFromReader(r.Body, &urlReq); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//Creating shortURL
	shortURL := generateShortUrl(urlReq.URL)

	//Creating new object in database
	urlRes := easyjson2.URLResponse{
		URL: "http://" + cfg.ShortURLAddr + "/" + shortURL,
	}
	if err := db.CreateUrl(shortURL, urlReq.URL); err != nil {
		http.Error(w, ErrMsgDBWriteError, http.StatusInternalServerError)
		return
	}

	//Configuring response's parameters
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	if _, err := easyjson.MarshalToWriter(urlRes, w); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func mainPageGetHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config, db db.DatabaseService) {

	//Reading short URL from URL's parameter
	shortURL := chi.URLParam(r, "shortURL")

	//Extracting URL object from database
	urlObj, err := db.GetUrlByShortName(shortURL)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//Check for existing value in object
	if urlObj.Original == "" {
		http.Error(w, "No elements satisfying conditions", http.StatusBadRequest)
		return
	}

	//Configuring response's parameters
	w.Header().Set("Location", urlObj.Original)
	w.WriteHeader(http.StatusTemporaryRedirect)
	return
	//fileHandler, err := models.NewFileHandler(cfg.FileStoragePath, os.O_RDONLY|os.O_CREATE)
	//if err != nil {
	//	http.Error(w, err.Error(), http.StatusInternalServerError)
	//}
	//for {
	//	record, err := fileHandler.ReadURLFromFile()
	//	if err != nil {
	//		if err == io.EOF {
	//			break
	//		}
	//	}
	//	if record == nil {
	//		w.WriteHeader(http.StatusBadRequest)
	//		return
	//	}
	//	if record.Short == reqShortURL {
	//		w.Header().Set("Location", record.Original)
	//		w.WriteHeader(http.StatusTemporaryRedirect)
	//		return
	//	}
	//
	//}
	//w.WriteHeader(http.StatusBadRequest)

}
func pingDBHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config, db db.DatabaseService) {
	err := db.Ping()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
	w.WriteHeader(http.StatusOK)
}
