package models

import (
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Waktu string
	Audio string
	DayID uint
}