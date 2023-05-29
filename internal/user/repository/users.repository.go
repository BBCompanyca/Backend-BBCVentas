package repository

import (
	"context"

	"github.com/BBCompanyca/Backend-BBCVentas.git/internal/user/entity"
)

const (
	qryInsertUser        = `insert into user (name, username, password, permission_level, status, date_register, date_update, registered_by) values (?,?,?,?,?,?,?,?)`
	qryGetUserByUsername = `select userID, name, username, permission_level, status, date_register, date_update, registered_by from user where username = ?`
	qryGetAllUser        = `select userID, name, username, permission_level, status, date_register, date_update, registered_by from user`
	qryUpdateUser        = `update set name = ?, username = ?, password = ?, permission_level = ?, status = ?, date_update = ? from user where userID = ?`
)

func (r *repo) SaveUser(ctx context.Context, name string, username string, password string, permissions string, status int, date_register string, date_update string, registered_by string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, name, username, password, permissions, status, date_register, date_update, registered_by)
	return err
}

func (r *repo) GetUserByUsername(ctx context.Context, username string) (*entity.User, error) {

	u := &entity.User{}

	err := r.db.GetContext(ctx, u, qryGetUserByUsername, username)
	if err != nil {
		return nil, err
	}

	return u, nil
}

func (r *repo) GetAllUser(ctx context.Context) ([]entity.User, error) {

	uu := []entity.User{}

	err := r.db.SelectContext(ctx, &uu, qryGetAllUser)
	if err != nil {
		return nil, err
	}

	return uu, nil
}

func (r *repo) updateUser(ctx context.Context, ID int64, name string, username string, password string, permission int, status int, date_update string) error {

	return nil
}
