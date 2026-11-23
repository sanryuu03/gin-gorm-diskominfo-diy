package domain

import (
	"time"

	"gorm.io/gorm"
)

type Bookings struct {
	ID         uint
	RoomID     uint   `gorm:"type:varchar(100)"`
	Rooms      Rooms  `gorm:"foreignKey:RoomID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	UserID     uint   `gorm:"type:varchar(255);not null"`
	Users      Users  `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Title      string `gorm:"type:varchar(150)"`
	StartTime  time.Time
	EndTime    time.Time
	Status     string `gorm:"type:varchar(20)"`
	CreatedAt  int64
	UpdatedAt  int64
	Deleted    bool
	SoftDelete int64
}

func (t *Bookings) BeforeCreate(tx *gorm.DB) (err error) {
	t.CreatedAt = time.Now().Unix()
	t.UpdatedAt = time.Now().Unix()
	return
}

func (t *Bookings) BeforeUpdate(tx *gorm.DB) (err error) {
	t.UpdatedAt = time.Now().Unix()
	return
}

func (t *Bookings) BeforeDelete(tx *gorm.DB) (err error) {
	t.SoftDelete = time.Now().Unix()
	return
}
