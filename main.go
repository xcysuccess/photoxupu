package main

import (
	order "photoxupu/dao"

	"github.com/gin-gonic/gin"
)

func main() {
	defer order.Close_databases()
	order.Connect_databases()
	order.Select_Users()
	//order.Insert_User()
	//order.Update_User()
	order.Delete_User()
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello golang!",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
