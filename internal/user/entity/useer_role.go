package entity

type UserRole struct {
	ID     int64 `db:"ID"`
	UserID int64 `db:"userID"`
	RolID  int64 `db:"rolID"`
}
