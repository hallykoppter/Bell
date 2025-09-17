package service

import "Bell/api/models"

func (a *App) GetSetting() ([]models.Setting, error) {
	var setting []models.Setting
	err := a.Db.First(&setting).Error
	if err != nil {
		return nil, err
	}
	return setting, nil
}