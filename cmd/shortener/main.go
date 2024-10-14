package main

import (
	"fmt"
	"github.com/MrPixik/url_shortener/internal/app/server"
	"github.com/MrPixik/url_shortener/internal/config"
	"net/http"
)

func main() {

	config := config.InitConfig()
	router := server.InitHandlers(config)

	fmt.Println("Starting server at " + config.LocalServerAddr)
	http.ListenAndServe(config.LocalServerAddr, router)
}
