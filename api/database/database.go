package database

import (
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

func InitializeDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{TranslateError: true})
	if err != nil {
		log.Fatal("Failed to connect to database " + err.Error())
	}
	return db
}