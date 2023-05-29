package repository

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/sign_in/entity"
)

var (
	qryGetUserByUsername = `select 
	userID, 
	name, 
	username, 
	password, 
	permission_level, 
	status, 
	date_register, 
	date_update, 
	registered_by
	from user where username = ?`
)

func (r *repo) Login(ctx context.Context, username string) (*entity.Sing_In, error) {

	u := &entity.Sing_In{}

	err := r.db.GetContext(ctx, u, qryGetUserByUsername, username)
	if err != nil {
		return nil, err
	}

	return u, nil

}
