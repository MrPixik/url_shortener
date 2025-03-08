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

	cfg, err := config.InitConfig()
	if err != nil {
		panic(err)
	}
	middleware.InitLogger()

	dbService := db.InitDBService(cfg, middleware.Logger)
	router := server.InitHandlers(cfg, middleware.Logger, dbService)

	fmt.Println("Starting server at " + cfg.LocalServerAddr)
	http.ListenAndServe(cfg.LocalServerAddr, router)
}
