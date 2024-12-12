package models

import (
	"time"
)

type Driveronlinelogs struct {
	Driverid         int
	SessionStartLoc  *GPoint    `gorm:"type:geography(POINT,4326)"`
	SessionStartTime *time.Time `gorm:"primaryKey"`
	SessionEndLoc    *GPoint    `gorm:"type:geography(POINT,4326)"`
	SessionEndTime   *time.Time
	SessionDate      *time.Time `gorm:"autoCreateTime"`
}
