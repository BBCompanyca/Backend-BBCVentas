package dto

type Sing_In struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
