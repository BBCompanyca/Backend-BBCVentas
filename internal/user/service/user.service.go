package service

import (
	"context"
	"errors"

	"github.com/BBCompanyca/Backend-BBCVentas.git/encryption"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/models"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
	invalidPassword      = errors.New("invalid password")
)

func (s *serv) RegisterUser(ctx context.Context, email, name, password string) error {

	u, _ := s.repo.GetUserByEmail(ctx, email)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, email, name, pass)
}

func (s *serv) Login(ctx context.Context, email, password string) (*models.User, error) {

	u, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	decryptedPassword, err := encryption.Decrypt(bb)

	if string(decryptedPassword) != password {
		return nil, invalidPassword
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}
