package repository

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/entity"
)

const (
	qryInsertUser     = `insert into user (email, name, password) values (?,?,?)`
	qryGeyUserByEmail = `select 
	userID, 
	email, 
	name, 
	password 
	from user
	where email = ?`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	u := &entity.User{}

	err := r.db.GetContext(ctx, u, qryGeyUserByEmail, email)

	return u, err
}
