package main

import (
	"flag"
	"fmt"
	"gin-gorm-diskominfo-diy/app"
	"gin-gorm-diskominfo-diy/controllers"
	"gin-gorm-diskominfo-diy/models/domain"
	"gin-gorm-diskominfo-diy/repositories"
	"gin-gorm-diskominfo-diy/services"

	"github.com/gin-gonic/gin"
)

func main() {
	migrate := flag.Bool("migrate", false, "Run database migrations")
	drop := flag.Bool("drop", false, "Drop database tables")
	seed := flag.Bool("seed", false, "Seed database tables")
	truncate := flag.Bool("truncate", false, "Truncate database tables")
	flag.Parse()

	db := app.OpenConnection()
	if *migrate {
		err := db.AutoMigrate(&domain.Users{}, &domain.Rooms{}, &domain.Bookings{})
		if err != nil {
			panic("Gagal melakukan migrasi database!")
		}
		fmt.Println("Database successfully migrated ⚡ 🚀 🚗")
		return
	}

	if *drop {
		db.Migrator().DropTable(&domain.Users{}, &domain.Rooms{}, &domain.Bookings{})
		fmt.Println("Tables successfully dropped ⚡ 🚀 🚗")
		return
	}

	if *seed {
		app.RunSeed(db)
		fmt.Println("Tables successfully seed ⚡ 🚀 🚗")
		return // Tambahkan ini untuk menghentikan eksekusi setelah seeding
	}

	if *truncate {
		app.RunTruncate(db)
		fmt.Println("Tables successfully truncated.")
		return
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	authController := controllers.NewAuthController(userService)

	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	v1 := router.Group("/api/v1")
	{
		v1.POST("/auth/login", authController.Login)

	}

	fmt.Println(db)

	router.Run() // listens on 0.0.0.0:8080 by default
}
