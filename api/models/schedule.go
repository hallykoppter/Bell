package models

import (
	"time"

	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Waktu time.Time
	Audio string
	DayID uint
}