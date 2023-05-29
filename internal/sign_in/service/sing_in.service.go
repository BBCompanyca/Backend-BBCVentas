package service

import (
	"context"
	"errors"

	"github.com/BBCompanyca/Backend-BBCVentas.git/encryption"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/models"
)

var (
	InvalidCredential = errors.New("invalid crendential")
)

func (s *serv) Login(ctx context.Context, username, password string) (*models.Sing_In, error) {

	u, err := s.repo.Login(ctx, username)
	if err != nil {
		return nil, err
	}

	bb, err := encryption.FromBase64(u.Password)
	if err != nil {
		return nil, err
	}

	passDecrypt, err := encryption.Decrypt(bb)
	if err != nil {
		return nil, err
	}

	if string(passDecrypt) != password {
		return nil, InvalidCredential
	}

	return &models.Sing_In{
		ID:            u.UserID,
		Name:          u.Name,
		Username:      u.Username,
		Permission:    traslatePermissionUser(u.Permission),
		Status:        traslateStatusUser(u.Status),
		Date_Register: u.Date_Register,
		Date_update:   u.Date_Update,
		Registered_By: u.Registered_By,
	}, nil
}

func traslatePermissionUser(permission int) string {

	var permissionStr = ""

	switch permission {
	case 1:
		permissionStr = "Adminitrador"
	case 2:
		permissionStr = "Vendedor"
	}

	return permissionStr
}

func traslateStatusUser(status int) string {

	var statusStr = ""

	switch status {
	case 0:
		statusStr = "Inactivo"
	case 2:
		statusStr = "Activo"
	}

	return statusStr
}
