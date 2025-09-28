package models

import (
	"gorm.io/gorm"
)

type Schedule struct {
	gorm.Model
	Waktu 	string
	BellID 	uint
	DayID 	uint
	Bell 	Bell
}