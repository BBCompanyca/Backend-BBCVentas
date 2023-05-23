package dtos

type RegisterUser struct {
	Name          string `json:"name" validate:"required"`
	Username      string `json:"username" validate:"required"`
	Password      string `json:"password" validate:"required,min=8"`
	Permissions   string `json:"permissions" validate:"required"`
	Registered_By string `json:"registered_by" validate:"required"`
}
