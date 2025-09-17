package repository

import (
	"Bell/api/models"

	"gorm.io/gorm"
)

func QuerySchedule(db *gorm.DB, DayID int8) ([]models.Schedule, error) {
	var schedule []models.Schedule
	err := db.Find(&schedule, "DayID = ?", DayID).Error
	if err != nil {
		return nil, err
	}
	return schedule, nil
}