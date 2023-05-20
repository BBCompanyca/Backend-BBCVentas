package service

import (
	"context"
	"errors"

	"github.com/BBCompanyca/Backend-BBCVentas.git/encryption"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/models"
)

var (
	InvalidPassword      = errors.New("invalid password")
	ErrUserAlreadyExists = errors.New("user already exists")
	ErrRoleAlreadyAdded  = errors.New("role was already added for this user")
	ErrNotFound          = errors.New("error not found")
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

	decryptedPassword, _ := encryption.Decrypt(bb)

	if string(decryptedPassword) != password {
		return nil, InvalidPassword
	}

	return &models.User{
		ID:    u.ID,
		Email: u.Email,
		Name:  u.Name,
	}, nil
}

func (s *serv) SaveUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRole(ctx, userID)
	if err != nil {
		return err
	}

	for _, r := range roles {
		if r.RolID == roleID {
			return ErrRoleAlreadyAdded
		}
	}

	return s.repo.SaveUserRole(ctx, userID, roleID)
}

func (s *serv) RemoveUserRole(ctx context.Context, userID, roleID int64) error {

	roles, err := s.repo.GetUserRole(ctx, userID)
	if err != nil {
		return err
	}

	roleFound := false

	for _, r := range roles {
		if r.RolID == roleID {
			roleFound = true
		}
	}

	if !roleFound {
		return ErrNotFound
	}

	return s.RemoveUserRole(ctx, userID, roleID)
}
