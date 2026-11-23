package app

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func OpenConnection() *gorm.DB {
	envFile := ".env"

	err := godotenv.Load(envFile)
	if err != nil {
		log.Printf("No %s file found, using system environment variables", envFile)
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")
	sslmode := os.Getenv("DB_SSLMODE")
	timezone := os.Getenv("DB_TIMEZONE")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s prepared_statement_cache_size=0", host, user, password, dbname, port, sslmode, timezone)

	dialect := postgres.Open(dsn)
	db, err := gorm.Open(dialect, &gorm.Config{})

	if err != nil {
		panic(err)
	}

	fmt.Println("success connect DB ⚡ 🚀 🚗 || Maintainer => https://muhammadhasan.my.id")

	return db
}
