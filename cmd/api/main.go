package main

import (
	"net/http"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/elkcityhazard/am-contact-form/internal/config"
	"github.com/elkcityhazard/am-contact-form/internal/handlers"
)

var app *config.AppConfig

func main() {

	app = config.NewAppConfig()

	app.SessionManager = NewSessionManager()

	repo := handlers.NewRepo(app, nil)

	handlers.NewHandlers(repo)

	var router = NewRouter()

	app.Router = router

	srv := &http.Server{
		Addr:           ":8080",
		Handler:        app.Router.InitRoutes(),
		MaxHeaderBytes: 2 >> 30,
		WriteTimeout:   15 * time.Second,
		ReadTimeout:    15 * time.Second,
		IdleTimeout:    30 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}

}

func NewSessionManager() *scs.SessionManager {
	sessionManager := scs.New()

	sessionManager.Lifetime = 24 * time.Hour

	return sessionManager
}
