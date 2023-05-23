package service

import (
	"context"
	"errors"

	"github.com/BBCompanyca/Backend-BBCVentas.git/encryption"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/models"
)

var (
	ErrUserAlreadyExists = errors.New("user already exists")
)

func (s *serv) SaveUser(ctx context.Context, name string, username string, password string, permissions string, status int, date_register string, date_update string, registered_by string) error {

	u, _ := s.repo.GetUserByUsername(ctx, username)
	if u != nil {
		return ErrUserAlreadyExists
	}

	bb, err := encryption.Encrypt([]byte(password))
	if err != nil {
		return err
	}

	pass := encryption.ToBase64(bb)

	return s.repo.SaveUser(ctx, name, username, pass, traslatePermissionUser(permissions), status, date_register, date_update, registered_by)
}

func (s *serv) GetAllUser(ctx context.Context) ([]models.User, error) {

	u, err := s.repo.GetAllUser(ctx)
	if err != nil {
		return nil, err
	}

	user := []models.User{}

	for _, u := range u {
		user = append(user, models.User{
			UserID:        u.UserID,
			Name:          u.Name,
			Username:      u.Username,
			Permission:    u.Permission,
			Status:        u.Status,
			Date_Register: u.Date_Register,
			Date_Update:   u.Date_Update,
			Registered_By: u.Registered_By,
		})
	}

	return user, nil
}

func traslatePermissionUser(permission string) string {

	var valuePermission string

	switch permission {
	case "Administrador":
		valuePermission = "1"
	case "Vendedor":
		valuePermission = "2"
	}
	return valuePermission
}
