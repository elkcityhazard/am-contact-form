package handlers

import (
	"github.com/elkcityhazard/am-contact-form/internal/config"
	"github.com/elkcityhazard/am-contact-form/internal/driver"
	"github.com/elkcityhazard/am-contact-form/internal/repository"
	"github.com/elkcityhazard/am-contact-form/internal/repository/dbrepo"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
	DB  repository.DatabaseInterface
}

func NewRepo(a *config.AppConfig, conn *driver.DB) *Repository {
	return &Repository{
		App: a,
		DB:  dbrepo.NewSqlDBRepo(a, conn),
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}
