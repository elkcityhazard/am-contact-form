package handlers

import (
	"database/sql"

	"github.com/elkcityhazard/am-contact-form/internal/config"
)

type Repo struct {
	App    *config.AppConfig
	DB     *sql.DB
	Router *Router
}

var HandlerRepo *Repo

func NewRepo(cfg *config.AppConfig, sqdb *sql.DB, rtr *Router) *Repo {
	return &Repo{
		App:    cfg,
		DB:     sqdb,
		Router: rtr,
	}
}

func NewHandlers(r *Repo) {
	HandlerRepo = r
}
