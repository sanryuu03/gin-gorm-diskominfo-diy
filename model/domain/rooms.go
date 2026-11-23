package domain

import (
	"time"

	"gorm.io/gorm"
)

type Rooms struct {
	ID         uint
	Name       string `gorm:"type:varchar(100)"`
	Capacity   int
	IsActive   bool
	CreatedAt  int64
	UpdatedAt  int64
	Deleted    bool
	SoftDelete int64
}

func (t *Rooms) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now().Unix()
	t.UpdatedAt = time.Now().Unix()
	return
}

func (t *Rooms) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now().Unix()
	return
}

func (t *Rooms) BeforeDelete(tx *gorm.DB) (err error) {
	t.SoftDelete = time.Now().Unix()
	return
}
