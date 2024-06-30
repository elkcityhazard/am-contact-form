package main

import (
	"net/http"
	"os"
	"time"

	"github.com/alexedwards/scs/v2"
	"github.com/elkcityhazard/am-contact-form/internal/config"
	"github.com/elkcityhazard/am-contact-form/internal/driver"
	"github.com/elkcityhazard/am-contact-form/internal/handlers"
)

var app *config.AppConfig

func main() {
	app = config.NewAppConfig()

	app.SessionManager = NewSessionManager()

	db, err := driver.ConnectSQL(os.Getenv("DSN"))
	if err != nil {
		panic(err)
	}

	repo := handlers.NewRepo(app, db)

	handlers.NewHandlers(repo)

	router := NewRouter()

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
