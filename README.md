# sqlx
## 一. 命令行安装
`go get github.com/jmoiron/sqlx`

## 二. model设计
```
package model

import "time"

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
```

## 三. 数据库增删改查
```
package order

import (
   "fmt"
   "sql-operator/model"
   "time"

   "github.com/jmoiron/sqlx"
)

var db_sqlx *sqlx.DB

func ConnectSqlxDatabases() {
   var err error
   db_sqlx, err = sqlx.Connect("mysql", "root:Xcy3329257@(localhost)/db_common?charset=utf8mb4&amp;parseTime=True&amp;loc=Local")
   if err != nil {
      fmt.Printf("connect DB failed, err:%v\n", err)
      return
   }
   println("sqlx数据库连接接成功")
}

func CloseSqlxDatabases() {
   db_sqlx.Close()
   println("sqlx数据库关闭成功")
}

func QueryUsers() {
   sqlStr := "SELECT * FROM db_common.users"
   var users []model.User_Sqlx

   err := db_sqlx.Select(&amp;users, sqlStr)
   if err != nil {
      fmt.Printf("get failed, err:%v\n", err)
      return
   }
   for i := 0; i < len(users); i++ {
      fmt.Printf("i=%d value=%v \n", i, users[i])
   }
}

// 查询单条数据示例
func QueryRowUser() {
   sqlStr := "select * from users where user_id=?"
   var user model.User_Sqlx
   err := db_sqlx.Get(&amp;user, sqlStr, 1)
   if err != nil {
      fmt.Printf("get failed, err:%v\n", err)
      return
   }
   fmt.Printf("user_id:%d user_name:%s birthday:%v \n", user.UserId, user.UserName, user.Birthday)
}

func InsertSqlxUser() {
   _, err := db_sqlx.Exec("INSERT INTO users(user_name,password,sex,birthday,last_login,bind_phone,reg_time) VALUES (?,?,?,?,?,?,?)", "tomxiang_sqlx", "tomxiang_pwd", 1, time.Now(), time.Now(), "15817307777", time.Now())
   if err != nil {
      fmt.Printf("get failed, err:%v\n", err)
      return
   }
   println("sqlx数据库插入成功!")

   println("sqlx插入之后的数据如下:")
   QueryUsers()
}

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
```

## 参考链接
1.[sqlx库使用指南](https://www.liwenzhou.com/posts/Go/sqlx/)
2.[[Mac 10.13.4] 使用Homebrew安装Mysql全过程](https://cloud.tencent.com/developer/article/1868895)

## 安装环境
`go get -u github.com/jinzhu/gorm`

# Gorm

## 一. model结构体
在线转换
[https://printlove.cn/tools/sql2gorm](https://printlove.cn/tools/sql2gorm)
model的结构：
```
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
```

## 二. 执行语句
```
package order

import (
   "database/sql"
   "fmt"
   "sql-operator/model"
   "time"

   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
)

var gorm_db *gorm.DB

func ConnectGormDatabases() {
   //"用户名:密码@[连接方式](主机名:端口号)/数据库名"
   var err error
   gorm_db, err = gorm.Open("mysql", "root:Xcy3329257@(localhost)/db_common?charset=utf8mb4&amp;parseTime=True&amp;loc=Local")
   if err != nil {
      fmt.Println("gorm数据库连接成功")
      panic(err)
   }
}

func CloseGormDatabases() {
   gorm_db.Close()
}

func SelectGormUser() {
   var users []model.User_Gorm
   // Get all records
   result := gorm_db.Find(&amp;users)
   println("查找到的result数目" + fmt.Sprintf("%d", result.RowsAffected))
   for i := 0; i < len(users); i++ {
      fmt.Printf("i=%d value=%v \n", i, users[i])
   }
}

func InsertGormUser() {
   //user := model.User{UserId: 13, UserName: "tomxiang10", Password: "admin10", Sex: 1, Birthday: sql.NullTime{time.Now(), true}, LastLogin: sql.NullTime{time.Now(), true}, BindPhone: "15817304444", RegTime: sql.NullTime{time.Now(), true}}
   user := model.User_Gorm{UserName: "tomxiang10", Password: "admin10", Sex: 1, Birthday: sql.NullTime{time.Now(), true}, LastLogin: sql.NullTime{time.Now(), true}, BindPhone: "15817304444", RegTime: sql.NullTime{time.Now(), true}}
   gorm_db.Create(&amp;user)
}

func DeleteGormUser() {
   var user model.User_Gorm
   gorm_db.Where("user_id = ?", 14).Take(&amp;user)
   gorm_db.Delete(user)
}

func UpdateGormUser() {
   var user model.User_Gorm
   //查找user_id=13
   gorm_db.Where("user_id = ?", 13).Take(&amp;user)
   //修改为当前时间
   user.Birthday = sql.NullTime{time.Now(), true}
   user.RegTime = sql.NullTime{time.Now(), true}
   user.LastLogin = sql.NullTime{time.Now(), true}
   gorm_db.Save(user)
}
```



## 参考链接
1. [gorm官网](https://gorm.io/zh_CN/docs/index.html)



