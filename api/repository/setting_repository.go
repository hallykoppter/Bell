package repository

import (
	"Bell/api/models"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

func FirstSeedRpository(db *gorm.DB) error {
	password, _ := bcrypt.GenerateFromPassword([]byte("admin"), 14)
	setting := []models.Setting{{NamaSekolah: "default", Npsn: "default", LogoSekolah: "default.jpg", Password: string(password), Active: 1}}
	err := db.Create(&setting).Error
	if err != nil {
		return err
	}
	return nil
}

func QuerySetting(db *gorm.DB) (*models.Setting, error) {
	var a *models.Setting
	err := db.First(&a).Error
	return a, err
}

func CheckPasswordHash(hash, password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func QueryUpdateSetting(db *gorm.DB, setting *models.Setting) error {
	return db.Model(models.Setting{}).Where("id = ?", 1).Updates(setting).Error
}

func QueryTruncateJadwaldanBel(db *gorm.DB) error {
	err := db.Exec("DELETE FROM 'bells'").Exec("DELETE FROM 'schedules'").Exec("UPDATE sqlite_sequence SET seq = 0 WHERE name = 'bells'").Exec("UPDATE sqlite_sequence SET seq = 0 WHERE name = 'schedules'").Exec("VACUUM").Error
	return err
}