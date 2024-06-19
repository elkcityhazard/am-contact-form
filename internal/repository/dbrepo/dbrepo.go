package dbrepo

import (
	"github.com/elkcityhazard/am-contact-form/internal/config"
	"github.com/elkcityhazard/am-contact-form/internal/driver"
	"github.com/elkcityhazard/am-contact-form/internal/repository"
)

type sqlDBRepo struct {
	App *config.AppConfig
	DB  *driver.DB
}

func NewSqlDBRepo(a *config.AppConfig, conn *driver.DB) repository.DatabaseInterface {
	return &sqlDBRepo{
		App: a,
		DB:  conn,
	}
}
