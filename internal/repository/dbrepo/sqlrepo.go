package dbrepo

func (s *sqlDBRepo) Ping() error {
	err := s.DB.Ping()

	if err != nil {
		return err
	}

	return nil

}
