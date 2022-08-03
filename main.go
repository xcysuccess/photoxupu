package main

import (
	"fmt"
	order "sql-operator/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	//gorm()
	sqlx()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello golang!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func gorm() {
	defer order.CloseGormDatabases()
	order.ConnectGormDatabases()
	order.SelectGormUser()
	//order.InsertGormUser()
	//order.UpdateGormUser()
	//order.DeleteGormUser()
}

func sqlx() {
	defer order.CloseSqlxDatabases()
	order.ConnectSqlxDatabases()
	fmt.Println("sqlx插入之前的数据如下:")
	order.QueryUsers()
	//order.InsertSqlxUser()
	//order.DeleteSqlxUser()
	//order.UpdateSqlxUser()
}
