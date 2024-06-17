package main

import (
	"net/http"
	"time"

	"github.com/elkcityhazard/am-contact-form/internal/config"
	"github.com/elkcityhazard/am-contact-form/internal/handlers"
)

func main() {

	app := config.NewAppConfig()

	repo := handlers.NewRepo(app, nil, handlers.NewRouter())

	handlers.NewHandlers(repo)

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        repo.Router.InitRoutes(),
		MaxHeaderBytes: 2 >> 30,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		IdleTimeout:    30 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}
