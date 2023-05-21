package dtos

type RegisterUser struct {
	Name          string `json:"name"`
	Username      string `json:"Username"`
	Password      string `json:"password"`
	Permissions   int    `json:"permissions"`
	Registered_By string `json:"registered_by"`
}
