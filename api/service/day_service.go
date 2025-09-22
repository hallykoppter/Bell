package service

import (
	"Bell/api/models"
	"Bell/api/repository"
)

func (a *App) GetDays() ([]models.Day, error) {
	days, err := repository.QueryAllDay(a.Db)
	return days, err
}

func (a *App) GetDayById(ID int8) (*models.Day, error) {
	day, err := repository.QueryDay(a.Db, ID)
	return day, err
}

