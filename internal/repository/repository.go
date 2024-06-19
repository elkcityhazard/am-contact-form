package repository

import "github.com/elkcityhazard/am-contact-form/internal/models"

type DatabaseInterface interface {
	Ping() error
	InsertMessage(msg *models.Message) (int64, error)
}
