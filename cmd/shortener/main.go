package main

import (
	"fmt"
	"github.com/MrPixik/url_shortener/internal/app/middleware"
	"github.com/MrPixik/url_shortener/internal/app/server"
	"github.com/MrPixik/url_shortener/internal/config"
	"net/http"
)

func main() {

	config.InitConfig()
	middleware.InitLogger()
	router := server.InitHandlers(config.Cfg, middleware.Logger)

	fmt.Println("Starting server at " + config.Cfg.LocalServerAddr)
	http.ListenAndServe(config.Cfg.LocalServerAddr, router)
}
