package models

import (
	"time"
)

type RidesInfo struct {
	RideID            uint   `gorm:"primaryKey"`
	PickupLoc         GPoint `gorm:"type:geography(POINT,4326)"`
	DropOffLoc        GPoint `gorm:"type:geography(POINT,4326)"`
	DriverLastLoc     GPoint `gorm:"type:geography(POINT,4326)"`
	RiderID           int
	DriverID          int
	InitPrice         float64
	ActualPrice       *float64
	InitDistance      *float64
	ActualDistance    *float64
	PickUpTime        *time.Time
	WaitTime          time.Duration
	RideDatetime      *time.Time
	RideStatusID      int
	InitDropOffTime   *time.Time
	ActualDropOffTime *time.Time
	OTP               string `gorm:"size:6"`
}
