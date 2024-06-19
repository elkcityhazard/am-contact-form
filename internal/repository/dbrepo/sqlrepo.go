package dbrepo

import (
	"context"
	"time"

	"github.com/elkcityhazard/am-contact-form/internal/models"
)

func (s *sqlDBRepo) Ping() error {
	err := s.DB.SQL.Ping()

	if err != nil {
		return err
	}

	return nil

}

func (s *sqlDBRepo) InsertMessage(msg *models.Message) (int64, error) {

	ctx, cancel := context.WithTimeout(s.App.Ctx, 15*time.Second)

	defer cancel()

	idChan := make(chan int64, 1)
	errorChan := make(chan error, 1)

	s.App.WG.Add(1)

	go func() {
		defer s.App.WG.Done()
		defer close(idChan)
		defer close(errorChan)

		stmt := `INSERT INTO Message (created_at, updated_at, version, name, email, message_content) VALUES(NOW(), NOW(), 1, ?,?,?)`

		result, err := s.DB.SQL.ExecContext(ctx, stmt, msg.Name, msg.Email, msg.MessageContent)

		if err != nil {
			errorChan <- err
			return
		}

		id, err := result.LastInsertId()

		if err != nil {
			errorChan <- err
			return
		}

		msg.ID = id

		idChan <- id

	}()

	err := <-errorChan

	if err != nil {
		return 0, err
	}

	id := <-idChan

	return id, nil

}
