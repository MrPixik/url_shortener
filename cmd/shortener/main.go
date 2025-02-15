package main

import (
	"fmt"
	"github.com/MrPixik/url_shortener/internal/app/middleware"
	"github.com/MrPixik/url_shortener/internal/app/server"
	"github.com/MrPixik/url_shortener/internal/config"
	"github.com/MrPixik/url_shortener/internal/db"
	"net/http"
)

func main() {

	config.InitConfig()
	middleware.InitLogger()
	dbService, err := db.InitDBService(config.Cfg)
	if err != nil {
		panic(err)
	}
	router := server.InitHandlers(config.Cfg, middleware.Logger, dbService)

	fmt.Println("Starting server at " + config.Cfg.LocalServerAddr)
	http.ListenAndServe(config.Cfg.LocalServerAddr, router)
}
