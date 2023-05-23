package service

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/models"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/repository"
)

type Service interface {
	SaveUser(ctx context.Context, name string, username string, password string, permissions int, status int, date_register string, date_update string, registered_by string) error
	GetAllUser(ctx context.Context) ([]models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
