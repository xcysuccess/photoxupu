package order

import (
	"fmt"
	"sql-operator/model"
	"time"

	"github.com/jmoiron/sqlx"
)

var db_sqlx *sqlx.DB

// ConnectSqlxDatabases TODO
func ConnectSqlxDatabases() {
	var err error
	db_sqlx, err = sqlx.Connect("mysql", "root:Xcy3329257@(localhost)/ginhello?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect DB failed, err:%v\n", err)
		return
	}
	println("sqlx数据库连接接成功")
}

// CloseSqlxDatabases TODO
func CloseSqlxDatabases() {
	db_sqlx.Close()
	println("sqlx数据库关闭成功")
}

// QueryUsers TODO
func QueryUsers() {
	sqlStr := "SELECT * FROM db_common.users"
	var users []model.User_Sqlx

	err := db_sqlx.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	for i := 0; i < len(users); i++ {
		fmt.Printf("i=%d value=%v \n", i, users[i])
	}
}

// QueryRowUser 查询单条数据示例
func QueryRowUser() {
	sqlStr := "select * from users where user_id=?"
	var user model.User_Sqlx
	err := db_sqlx.Get(&user, sqlStr, 1)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Printf("user_id:%d user_name:%s birthday:%v \n", user.UserId, user.UserName, user.Birthday)
}

// InsertSqlxUser TODO
func InsertSqlxUser() {
	_, err := db_sqlx.Exec(
		"INSERT INTO users(user_name,password,sex,birthday,last_login,bind_phone,reg_time) VALUES (?,?,?,?,?,?,?)", "tomxiang_sqlx", "tomxiang_pwd", 1, time.Now(), time.Now(), "15817307777", time.Now())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	println("sqlx数据库插入成功!")

	println("sqlx插入之后的数据如下:")
	QueryUsers()
}

// DeleteSqlxUser TODO
func DeleteSqlxUser() {
	_, err := db_sqlx.Exec("delete from users where user_id = ?", 17)
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	println("sqlx数据库删除成功!")

	println("sqlx删除之后的数据如下:")
	QueryUsers()
}

// UpdateSqlxUser TODO
func UpdateSqlxUser() {
	_, err := db_sqlx.Exec("update users set birthday = ? where user_id = 3", time.Now())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	println("sqlx数据库更新成功!")

	println("sqlx更新之后的数据如下:")
	QueryUsers()
}
