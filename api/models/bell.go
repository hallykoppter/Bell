package models

import "gorm.io/gorm"

type Bell struct {
	gorm.Model
	NamaBell	string `gorm:"unique"`
	SoundFile	string
}