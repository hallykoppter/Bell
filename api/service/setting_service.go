package service

import (
	"Bell/api/models"
	"Bell/api/repository"

	"golang.org/x/crypto/bcrypt"
)

func (a *App) GetSetting() (*models.Setting, error) {
	setting, err := repository.QuerySetting(a.Db)
	if err != nil {
		return nil, err
	}
	return setting, nil
}

func (a *App) UpdateSetting(NamaSekolah, Npsn string ) error {
	setting := &models.Setting{
		NamaSekolah: NamaSekolah,
		Npsn: Npsn,
	}
	err := repository.QueryUpdateSetting(a.Db, setting)
	return err
}

func (a *App) UpdateActiveSetting(Active int8) error {
	setting := &models.Setting{
		Active: Active,
	}
	err := repository.QueryUpdateSetting(a.Db, setting)
	return err
}

func (a *App) UpdateDayActive(DayActive int8) error {
	err := repository.QueryUpdateDayActive(a.Db, DayActive)
	return err
}

func (a *App) UpdatePassword(Pass string) error {
	password, _ := bcrypt.GenerateFromPassword([]byte(Pass), 14)
	setting := &models.Setting{
		Password: string(password),
	}
	err := repository.QueryUpdateSetting(a.Db, setting)
	return err
}

func (a *App) TruncateTabelJadwaldanBel() error {
	err := repository.QueryTruncateJadwaldanBel(a.Db)
	return err
}