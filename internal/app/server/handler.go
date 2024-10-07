package server

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"net/http"
)

var URLPool = make(map[string]string)

func MainPageHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:

		body, err := io.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		URL := string(body)
		hasher := md5.New()
		shortURL := "/" + hex.EncodeToString(hasher.Sum(body)[:4])
		URLPool[shortURL] = URL

		w.WriteHeader(201)
		w.Header().Set("Content-Type", "text/plain")
		w.Write([]byte("http://localhost:8080" + shortURL))

	case http.MethodGet:
		originalURL, ok := URLPool[r.URL.Path]
		if !ok {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.Header().Set("Location", originalURL)
		w.WriteHeader(307)
	}
}
