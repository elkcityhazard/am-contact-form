package config

import (
	"github.com/alexedwards/scs/v2"
	"github.com/elkcityhazard/am-contact-form/internal/repository"
)

type AppConfig struct {
	IsProduction   bool
	Port           string
	Router         repository.RouterInterface
	SessionManager *scs.SessionManager
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}
