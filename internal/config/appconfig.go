package config

type AppConfig struct {
	IsProduction bool
	Port         string
}

func NewAppConfig() *AppConfig {
	return &AppConfig{}
}
