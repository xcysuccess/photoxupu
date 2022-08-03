package model

import (
	"database/sql"
)

type User_Gorm struct {
	UserId    int          `gorm:"column:user_id;type:int(11);primary_key" json:"user_id"`
	UserName  string       `gorm:"column:user_name;type:varchar(60);NOT NULL" json:"user_name"`
	Password  string       `gorm:"column:password;type:varchar(32)" json:"password"`
	Sex       int          `gorm:"column:sex;type:int(11)" json:"sex"`
	Birthday  sql.NullTime `gorm:"column:birthday;type:timestamp" json:"birthday"`
	LastLogin sql.NullTime `gorm:"column:last_login;type:timestamp" json:"last_login"`
	BindPhone string       `gorm:"column:bind_phone;type:varchar(20)" json:"bind_phone"`
	RegTime   sql.NullTime `gorm:"column:reg_time;type:timestamp" json:"reg_time"`
}
