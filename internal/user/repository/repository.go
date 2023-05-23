package repository

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/entity"
	"github.com/jmoiron/sqlx"
)

type Repository interface {
	SaveUser(ctx context.Context, name string, username string, password string, permissions int, status int, date_register string, registered_by string) error
	GetUserByUsername(ctx context.Context, username string) (*entity.User, error)
	GetAllUser(ctx context.Context) ([]entity.User, error)
}

type repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &repo{
		db: db,
	}
}
