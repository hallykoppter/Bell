package service

import (
	"Bell/api/models"
	"Bell/api/repository"
	"errors"

	"gorm.io/gorm"
)

func (a *App) Migrate() error {
	// User Migration
	err := a.Db.AutoMigrate(&models.Day{}, &models.Setting{}, &models.Schedule{}, &models.Bell{})

	// User Seed
	if err == nil && a.Db.Migrator().HasTable(&models.Day{}) {
		if exist := a.Db.First(&models.Day{}).Error; errors.Is(exist, gorm.ErrRecordNotFound) {
			seed := repository.SeedDay(a.Db)
			if seed != nil {
				return seed
			}
		}
	}

	// Setting Seed
	if err == nil && a.Db.Migrator().HasTable(&models.Setting{}) {
		if exist := a.Db.First(&models.Setting{}).Error; errors.Is(exist, gorm.ErrRecordNotFound) {
			seed := repository.FirstSeedRpository(a.Db)
			if seed != nil {
				return seed
			}
		}
	}


	return nil
}