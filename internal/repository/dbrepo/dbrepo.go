package dbrepo

import (
	"database/sql"
	"net/http"

	"github.com/elkcityhazard/am-contact-form/internal/config"
	"github.com/elkcityhazard/am-contact-form/internal/repository"
)

type sqlDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

func NewSqlDBRepo(a *config.AppConfig, conn *sql.DB, rtr http.Handler) repository.DatabaseInterface {
	return &sqlDBRepo{
		App: a,
		DB:  conn,
	}
}
