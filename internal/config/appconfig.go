package config

import (
	"context"
	"sync"

	"github.com/alexedwards/scs/v2"
	"github.com/elkcityhazard/am-contact-form/internal/repository"
)

type AppConfig struct {
	IsProduction   bool
	Port           string
	Router         repository.RouterInterface
	SessionManager *scs.SessionManager
	Ctx            context.Context
	WG             *sync.WaitGroup
	MU             *sync.Mutex
}

func NewAppConfig() *AppConfig {
	return &AppConfig{
		SessionManager: scs.New(),
		Ctx:            context.Background(),
		WG:             &sync.WaitGroup{},
		MU:             &sync.Mutex{},
	}
}
