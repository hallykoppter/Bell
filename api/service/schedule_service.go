package service

import (
	"Bell/api/models"
	"Bell/api/repository"
)

func (a *App) InsertSchedule(Waktu, Audio string, DayID int) int {
	schedule := models.Schedule{
		Waktu: Waktu,
		Audio: Audio,
		DayID: uint(DayID),
	}
	res := repository.QueryInsertSchedule(a.Db, &schedule)
	return res
}

func (a *App) GetScheduleByDay(ID int) ([]models.Schedule, error) {
	res, err := repository.QueryScheduleByDay(a.Db, ID)
	return res, err
}

func (a *App) DeleteScheduleByID(ID int) error {
	err := repository.QueryScheduleDelete(a.Db, ID)
	return err
}