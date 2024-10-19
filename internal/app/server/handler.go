package server

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/MrPixik/url_shortener/internal/app/middleware"
	"github.com/MrPixik/url_shortener/internal/app/models"
	"github.com/MrPixik/url_shortener/internal/config"
	"github.com/go-chi/chi/v5"
	"github.com/mailru/easyjson"
	"go.uber.org/zap"
	"io"
	"net/http"
)

var uRLPool = make(map[string]string)

func generateShortUrl(longUrl string) string {
	hasher := md5.New()
	shortURL := hex.EncodeToString(hasher.Sum([]byte(longUrl))[0:12])
	uRLPool[shortURL] = longUrl
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

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://" + cfg.ShortURLAddr + "/" + shortURL))
}

func shortenURLPostHandler(w http.ResponseWriter, r *http.Request, cfg *config.Config) {

	var urlReq models.URLRequest
	err := easyjson.UnmarshalFromReader(r.Body, &urlReq)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//fmt.Println("Received URL:", URL)
	shortURL := generateShortUrl(urlReq.URL)

	urlRes := models.URLResponse{
		URL: "http://" + cfg.ShortURLAddr + "/" + shortURL,
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
	originalURL, ok := uRLPool[chi.URLParam(r, "id")]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
