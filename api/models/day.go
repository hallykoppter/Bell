package models

import (
	"gorm.io/gorm"
)

type Day struct {
	gorm.Model
	Day string `gorm:"unique;not null"`
	Schedule []Schedule
}