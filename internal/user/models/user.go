package models

type User struct {
	UserID        int64  `json:"ID"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Permission    int    `json:"permission"`
	Status        int    `json:"status"`
	Date_Register string `json:"register"`
	Date_Update   string `json:"update"`
	Registered_By string `json:"registered"`
}
