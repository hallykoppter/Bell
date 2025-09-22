package service

import (
	"Bell/api/models"
	"Bell/api/repository"
)

func (a *App) GetSetting() (*models.Setting, error) {
	setting, err := repository.QuerySetting(a.Db)
	if err != nil {
		return nil, err
	}
	return setting, nil
}