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

func QueryInsertSchedule(db *gorm.DB, schedule *models.Schedule) int {

	existingSchedule := models.Schedule{}

	searchCondition := models.Schedule{
		Waktu: schedule.Waktu,
		DayID: schedule.DayID,
	}

	db.Where(searchCondition).FirstOrInit(&existingSchedule)
	existingSchedule.Audio = schedule.Audio
	res := db.Save(&existingSchedule)

	return int(res.RowsAffected)

}

func QueryScheduleByDay(db *gorm.DB, ID int) ([]models.Schedule, error) {
	var Schedule []models.Schedule
	err := db.Where("day_id = ?", ID).Find(&Schedule).Error
	if err != nil {
		return nil, err
	}
	return Schedule, nil
}

func QueryScheduleDelete(db *gorm.DB, ID int) error {
	var Schedule *models.Schedule
	err := db.Unscoped().Delete(&Schedule, ID).Error
	return err
}

