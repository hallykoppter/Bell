package service

import (
	"Bell/api/models"
	"Bell/api/repository"
)

func (a *App) InsertBell(NamaBell, SoundFile string) error {
	Bell := &models.Bell{NamaBell: NamaBell, SoundFile: SoundFile}
	err := repository.QueryCreateBell(a.Db, Bell)
	return err
}

func (a *App) GetAllBell() ([]models.Bell, error) {
	res, err := repository.QueryGetAllBell(a.Db)
	return res, err
}

func (a *App) DeleteBellById(ID int) error {
	err := repository.QueryDeleteBellByID(a.Db, ID)
	return err
}