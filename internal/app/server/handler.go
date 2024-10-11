package server

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
)

var URLPool = make(map[string]string)

func MainPagePostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	URL := string(body)
	hasher := md5.New()
	shortURL := "/" + hex.EncodeToString(hasher.Sum(body)[:4])
	URLPool[shortURL] = URL

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:8080" + shortURL))
}

func MainPageGetHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	originalURL, ok := URLPool[r.URL.Path]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
