package service

import (
	"Bell/api/audio"
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

func (a *App) GetBellByID(ID int) (*models.Bell, error) {
	bell, err := repository.QueryGetBellByID(a.Db, ID)
	if err != nil {
		return nil, err
	}
	return bell, nil
}

func (a *App) PlayBell(soundFile string) error {
	err := audio.PlayBellSound(soundFile)
	return err
}