package handlers

import (
	"database/sql"

	"github.com/elkcityhazard/am-contact-form/internal/config"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewRepo(a *config.AppConfig, conn *sql.DB) *Repository {
	return &Repository{
		App: a,
		DB:  conn,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
