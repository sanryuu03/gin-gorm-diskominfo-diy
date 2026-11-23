package app

import (
	"log"

	"gorm.io/gorm"
)

func RunTruncate(db *gorm.DB) {
	if err := db.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE").Error; err != nil {
		log.Fatalf("Failed to truncate table: %v", err)
	}
	if err := db.Exec("TRUNCATE TABLE rooms RESTART IDENTITY CASCADE").Error; err != nil {
		log.Fatalf("Failed to truncate table: %v", err)
	}
	if err := db.Exec("TRUNCATE TABLE bookings RESTART IDENTITY CASCADE").Error; err != nil {
		log.Fatalf("Failed to truncate table: %v", err)
	}

}
