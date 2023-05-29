package entity

type Sing_In struct {
	UserID        int64  `db:"userID"`
	Name          string `db:"name"`
	Username      string `db:"username"`
	Password      string `db:"password"`
	Permission    int    `db:"permission_level"`
	Status        int    `db:"status"`
	Date_Register string `db:"date_register"`
	Date_Update   string `db:"date_update"`
	Registered_By string `db:"registered_by"`
}
