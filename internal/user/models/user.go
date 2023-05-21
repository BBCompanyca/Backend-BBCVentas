package models

type User struct {
	UserID        int64  `json:"userID"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Permissions   int    `json:"permission_level"`
	Status        int    `json:"status"`
	Date_Register string `json:"date_register"`
	Date_Update   string `json:"date_update"`
	Registered_By string `json:"registered_by"`
}
