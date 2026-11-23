package main

import (
	"flag"
	"fmt"
	"gin-gorm-diskominfo-diy/app"
	"gin-gorm-diskominfo-diy/model/domain"

	"github.com/gin-gonic/gin"
)

func main() {
	migrate := flag.Bool("migrate", false, "Run database migrations")
	flag.Parse()

	db := app.OpenConnection()
	if *migrate {
		err := db.AutoMigrate(&domain.Users{})
		if err != nil {
			panic("Gagal melakukan migrasi database!")
		}
		fmt.Println("Database successfully migrated ⚡ 🚀 🚗")
		return
	}

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	fmt.Println(db)

	router.Run() // listens on 0.0.0.0:8080 by default
}
