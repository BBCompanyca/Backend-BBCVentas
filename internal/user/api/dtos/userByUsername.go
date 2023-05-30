package dtos

type UserByUsernameDTO struct {
	Username string `json:"username" validate:"required"`
}
