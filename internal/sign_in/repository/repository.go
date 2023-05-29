package repository

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	Login(ctx context.Context, username string) (*entity.Sing_In, error)
}

type repo struct {
	db *sqlx.DB
}

func NewSign_in(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
