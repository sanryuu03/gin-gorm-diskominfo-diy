package domain

import (
	"time"

	"gorm.io/gorm"
)

type Users struct {
	ID           uint
	Name         string  `gorm:"type:varchar(100)"`
	Email        *string `gorm:"unique;type:varchar(100)"`
	PasswordHash string  `gorm:"type:varchar(255)"`
	Role         string  `gorm:"type:varchar(20)"`
	CreatedAt    int64
	UpdatedAt    int64
	Deleted      bool
	SoftDelete   int64
}

func (t *Users) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now().Unix()
	t.UpdatedAt = time.Now().Unix()
	return
}

func (t *Users) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now().Unix()
	return
}

func (t *Users) BeforeDelete(tx *gorm.DB) (err error) {
	t.SoftDelete = time.Now().Unix()
	return
}
