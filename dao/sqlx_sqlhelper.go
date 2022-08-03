package order

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

var sqlx_db *sqlx.DB

func Connect_databases() {
	var err error
	sqlx_db, err = sqlx.Connect("mysql", "root:Xcy3329257@(localhost)/db_common?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
}
