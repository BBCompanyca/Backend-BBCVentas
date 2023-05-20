package repository

import (
	"context"
	"fmt"

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

	qryInsertUserRole = `insert into user_roles (userID, rolID) values (:userID,rolID)`
	qryRemoveUserRole = `delete from user_roles where userID: userID and rolID = :rolID`
)

func (r *repo) SaveUser(ctx context.Context, email, name, password string) error {
	_, err := r.db.ExecContext(ctx, qryInsertUser, email, name, password)
	return err
}

func (r *repo) GetUserByEmail(ctx context.Context, email string) (*entity.User, error) {

	u := &entity.User{}
	err := r.db.GetContext(ctx, u, qryGeyUserByEmail, email)
	if err != nil {
		return nil, err
	}

	fmt.Println("pasó por aquí...")

	return u, nil
}

func (r *repo) SaveUserRole(ctx context.Context, userID, roleID int64) error {

	data := entity.UserRole{
		UserID: userID,
		RolID:  roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryInsertUserRole, data)

	return err
}

func (r *repo) RemoveUserRole(ctx context.Context, userID, roleID int64) error {

	data := entity.UserRole{
		UserID: userID,
		RolID:  roleID,
	}

	_, err := r.db.NamedExecContext(ctx, qryRemoveUserRole, data)

	return err
}

func (r *repo) GetUserRole(ctx context.Context, userID int64) ([]entity.UserRole, error) {

	role := []entity.UserRole{}

	err := r.db.SelectContext(ctx, &role, `select userID, roleID from user_role where userID = ?`, userID)
	if err != nil {
		return nil, err
	}

	return role, nil
}
