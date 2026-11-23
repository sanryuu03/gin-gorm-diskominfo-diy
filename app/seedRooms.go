package app

import (
	"gin-gorm-diskominfo-diy/helper"
	"gin-gorm-diskominfo-diy/models/domain"
	"time"

	"gorm.io/gorm"
)

func parseBookingTime(value string) time.Time {
	t, err := time.ParseInLocation("2006-01-02 15:04:05", value, time.Local)
	helper.PanicIfError(err)
	return t
}

func createBooking(db *gorm.DB, booking domain.Bookings) {
	err := db.Create(&booking).Error
	helper.PanicIfError(err)
}

func seedBookings(db *gorm.DB) {
	bookings := []domain.Bookings{
		{
			RoomID: 1, UserID: 2, Title: "Rapat Mingguan",
			StartTime: parseBookingTime("2026-11-03 09:00:00"),
			EndTime:   parseBookingTime("2026-11-03 11:00:00"),
			Status:    "CONFIRMED",
		},
		{
			RoomID: 1, UserID: 3, Title: "Review Kinerja",
			StartTime: parseBookingTime("2026-11-10 13:00:00"),
			EndTime:   parseBookingTime("2026-11-10 14:30:00"),
			Status:    "CONFIRMED",
		},

		{
			RoomID: 1, UserID: 2, Title: "Rapat Desember",
			StartTime: parseBookingTime("2026-12-01 10:00:00"),
			EndTime:   parseBookingTime("2026-12-01 12:00:00"),
			Status:    "CONFIRMED",
		},

		{
			RoomID: 2, UserID: 2, Title: "Sync Tim",
			StartTime: parseBookingTime("2026-11-05 10:00:00"),
			EndTime:   parseBookingTime("2026-11-05 12:00:00"),
			Status:    "CONFIRMED",
		},
		{
			RoomID: 2, UserID: 3, Title: "Booking Menunggu",
			StartTime: parseBookingTime("2026-11-18 08:00:00"),
			EndTime:   parseBookingTime("2026-11-18 09:00:00"),
			Status:    "PENDING",
		},

		{
			RoomID: 3, UserID: 1, Title: "Pelatihan Golang",
			StartTime: parseBookingTime("2026-11-20 09:00:00"),
			EndTime:   parseBookingTime("2026-11-20 15:00:00"),
			Status:    "CONFIRMED",
		},
		{
			RoomID: 3, UserID: 3, Title: "Booking Dibatalkan",
			StartTime: parseBookingTime("2026-11-22 16:00:00"),
			EndTime:   parseBookingTime("2026-11-22 17:00:00"),
			Status:    "CANCELLED",
		},

		{
			RoomID: 4, UserID: 2, Title: "Verifikasi Arsip",
			StartTime: parseBookingTime("2026-11-15 13:00:00"),
			EndTime:   parseBookingTime("2026-11-15 14:00:00"),
			Status:    "CONFIRMED",
		},
	}

	for _, booking := range bookings {
		createBooking(db, booking)
	}
}
