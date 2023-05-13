package service

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/models"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/repository"
)

type Service interface {
	RegisterUser(ctx context.Context, email, name, password string) error
	Login(ctx context.Context, email, password string) (*models.User, error)
}

type serv struct {
	repo repository.Repository
}

func New(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
