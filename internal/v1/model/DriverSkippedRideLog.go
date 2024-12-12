package models

import (
	"time"
)

type DriverSkippedRideLog struct {
	DriverID        int `gorm:"primaryKey"`
	RideID          int `gorm:"primaryKey"`
	SkippedDateTime *time.Time
}
