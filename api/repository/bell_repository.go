package repository

import (
	"Bell/api/models"
	"errors"

	"gorm.io/gorm"
)

func QueryCreateBell(db *gorm.DB, bell *models.Bell) error {
	res := db.Create(bell)
	err := res.Error

	if err != nil {
		if errors.Is(err, gorm.ErrDuplicatedKey){
			return errors.New("nama bel sudah ada")
		} 
		return err
	}
	return err
}

func QueryGetAllBell(db *gorm.DB) ([]models.Bell, error) {
	var Bell []models.Bell
	err := db.Find(&Bell).Error
	return Bell, err
}

func QueryDeleteBellByID(db *gorm.DB, ID int) error {
	var Bell *models.Bell
	err := db.Unscoped().Where("id = ?", ID).Delete(&Bell).Error
	return err
} 