package service

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/models"
	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/repository"
)

type Service interface {
	Login(ctx context.Context, username, password string) (*models.Sing_In, error)
}

type serv struct {
	repo repository.Repository
}

func NewSign_in(repo repository.Repository) Service {
	return &serv{
		repo: repo,
	}
}
