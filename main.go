package main

import (
	"fmt"
	"gin-gorm-diskominfo-diy/app"

	"github.com/gin-gonic/gin"
)

func main() {
	db := app.OpenConnection()

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Println(db)

	router.Run() // listens on 0.0.0.0:8080 by default
}
