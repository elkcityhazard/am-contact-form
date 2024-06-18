package repository

type DatabaseInterface interface {
	Ping() error
}
