package service

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/repository"
)

type Service interface {
	SaveUser(ctx context.Context, name string, username string, password string, permissions int, status int, date_register string, registered_by string) error
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
