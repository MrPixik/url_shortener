package main

import (
	"fmt"
	"github.com/MrPixik/url_shortener/internal/app/server"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func main() {

	router := chi.NewRouter()

	router.Route("/", func(r chi.Router) {
		r.Get("/{id}", server.MainPageGetHandler)
		r.Post("/", server.MainPagePostHandler)
	})

	fmt.Println("Starting server at port 8080...")
	http.ListenAndServe(":8080", router)
}
