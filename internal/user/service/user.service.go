package service

import (
	"context"
	"errors"

	"github.com/BBCompanyca/Backend-BBCVentas.git/encryption"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

func (s *serv) SaveUser(ctx context.Context, name string, username string, password string, permissions int, status int, date_register string, registered_by string) error {

	u, _ := s.repo.GetUserByUsername(ctx, username)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, name, username, pass, permissions, status, date_register, registered_by)
}
