package order

import (
	"database/sql"
	"fmt"
	"photoxupu/model"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var gorm_db *gorm.DB

func Connect_databases() {
	//"用户名:密码@[连接方式](主机名:端口号)/数据库名"
	var err error
	gorm_db, err = gorm.Open("mysql", "root:Xcy3329257@(localhost)/db_common?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("数据库连接成功")
		panic(err)
	}
}

func Close_databases() {
	gorm_db.Close()
}

func Select_Users() {
	var users []model.User
	// Get all records
	result := gorm_db.Find(&users)
	println("查找到的result数目" + fmt.Sprintf("%d", result.RowsAffected))
	for i := 0; i < len(users); i++ {
		fmt.Printf("i=%d value=%v \n", i, users[i])
	}
}

func Insert_User() {
	//user := model.User{UserId: 13, UserName: "tomxiang10", Password: "admin10", Sex: 1, Birthday: sql.NullTime{time.Now(), true}, LastLogin: sql.NullTime{time.Now(), true}, BindPhone: "15817304444", RegTime: sql.NullTime{time.Now(), true}}
	user := model.User{UserName: "tomxiang10", Password: "admin10", Sex: 1, Birthday: sql.NullTime{time.Now(), true}, LastLogin: sql.NullTime{time.Now(), true}, BindPhone: "15817304444", RegTime: sql.NullTime{time.Now(), true}}
	gorm_db.Create(&user)
}

func Delete_User() {
	var user model.User
	gorm_db.Where("user_id = ?", 14).Take(&user)
	gorm_db.Delete(user)
}

func Update_User() {
	var user model.User
	//查找user_id=13
	gorm_db.Where("user_id = ?", 13).Take(&user)
	//修改为当前时间
	user.Birthday = sql.NullTime{time.Now(), true}
	user.RegTime = sql.NullTime{time.Now(), true}
	user.LastLogin = sql.NullTime{time.Now(), true}
	gorm_db.Save(user)
}
