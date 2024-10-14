package server

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/go-chi/chi/v5"
	"io"
	"net/http"
)

var URLPool = make(map[string]string)

func MainPagePostHandler(w http.ResponseWriter, r *http.Request) {

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
	URLPool[shortURL] = URL
	//fmt.Println("\"" + "http://localhost:8080/" + shortURL + "\"" + " post")

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("http://localhost:8080/" + shortURL))
}

func MainPageGetHandler(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("\"" + chi.URLParam(r, "id") + "\"" + " get")
	originalURL, ok := URLPool[chi.URLParam(r, "id")]
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Location", originalURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
