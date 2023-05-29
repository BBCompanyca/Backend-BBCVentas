package models

type Sing_In struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Username      string `json:"username"`
	Permission    string `json:"permission"`
	Status        string `json:"status"`
	Date_Register string `json:"date_register"`
	Date_update   string `json:"date_update"`
	Registered_By string `json:"registered_by"`
}
