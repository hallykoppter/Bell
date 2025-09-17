package models

import (
	"gorm.io/gorm"
)

type Setting struct {
	gorm.Model
	NamaSekolah string `gorm:"type:varchar(100)"`
	Npsn        string `gorm:"type:varchar(100)"`
	LogoSekolah string `gorm:"type:varchar(100)"`
	Password	string `gorm:"type:varchar(100)"`
	Active      int8
}

