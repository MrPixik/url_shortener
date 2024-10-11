package main

import (
	"fmt"
	"github.com/MrPixik/url_shortener/internal/app/server"
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			server.MainPagePostHandler(w, r)
		} else if r.Method == http.MethodGet {
			server.MainPageGetHandler(w, r)
		}
	})

	fmt.Println("Starting server at port 8080...")
	http.ListenAndServe(":8080", nil)
}
