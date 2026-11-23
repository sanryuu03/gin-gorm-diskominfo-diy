package app

import (
	"gin-gorm-diskominfo-diy/helper"
	"gin-gorm-diskominfo-diy/model/domain"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func createUser(db *gorm.DB, user domain.Users, plainPassword string) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.DefaultCost)
	helper.PanicIfError(err)

	user.PasswordHash = string(hashedPassword)
	err = db.Create(&user).Error
	helper.PanicIfError(err)
}

func seedUsers(db *gorm.DB) {
	now := time.Now()
	unix := now.Unix()

	// Slice data user untuk di-loop
	users := []struct {
		name          string
		email         string
		plainPassword string
		role          string
	}{
		{
			name:          "Administrator",
			email:         "admin@test.com",
			plainPassword: "password123",
			role:          "ADMIN",
		},
		{
			name:          "User Satu",
			email:         "user1@test.com",
			plainPassword: "password123",
			role:          "USER",
		},
		{
			name:          "User Dua",
			email:         "user2@test.com",
			plainPassword: "password123",
			role:          "USER",
		},
	}

	// Loop untuk membuat user
	for _, u := range users {
		createUser(db, domain.Users{
			Name:      u.name,
			Email:     &u.email,
			Role:      u.role,
			CreatedAt: unix,
			UpdatedAt: unix,
		}, u.plainPassword)
	}
}

func RunSeed(db *gorm.DB) {
	seedUsers(db)
}
