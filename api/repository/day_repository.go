package repository

import (
	"Bell/api/models"

	"gorm.io/gorm"
)

func QueryAllDay(db *gorm.DB) ([]models.Day, error) {
	var days []models.Day
	err := db.Find(&days).Error
	return days, err
}

func QueryDay(db *gorm.DB, ID int8) (*models.Day, error) {
	var day *models.Day
	err := db.Find(&day, ID).Error
	return day, err
}

func SeedDay(db *gorm.DB) error {
	days := []models.Day{{Day: "Senin"}, {Day: "Selasa"}, {Day: "Rabu"}, {Day: "Kamis"}, {Day: "Jumat"}, {Day: "Sabtu"}}
	err := db.Create(&days).Error
	if err != nil {
		return err
	}
	return nil
}