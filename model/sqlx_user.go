package model

import "time"

// User_Sqlx TODO
type User_Sqlx struct {
	UserId    int       `db:"user_id"`
	UserName  string    `db:"user_name"`
	Password  string    `db:"password"`
	Sex       int       `db:"sex"`
	Birthday  time.Time `db:"birthday"`
	LastLogin time.Time `db:"last_login"`
	BindPhone string    `db:"bind_phone"`
	RegTime   time.Time `db:"reg_time"`
}
