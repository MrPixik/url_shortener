package main

import (
	"fmt"
	"github.com/MrPixik/url_shortener/internal/app/server"
	"net/http"
)

func main() {

	http.Handle("/", http.HandlerFunc(server.MainPageHandler))

	fmt.Println("Starting server at port 8080...")
	http.ListenAndServe(":8080", nil)
}
