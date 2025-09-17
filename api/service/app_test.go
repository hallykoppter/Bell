package service

import (
	"Bell/api/models"
	"fmt"
	"testing"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

// func TestConnection(t *testing.T) {
// 	a := InitDatabase()
// 	fmt.Println(a)
// }

func TestConnectonDua(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("../database/database.db"), &gorm.Config{})
	if err != nil {
		fmt.Print("Failed " + err.Error())
	}
	fmt.Print("Success")
	fmt.Print(db)
}

func TestCompatePassword(t *testing.T) {
	db, err := gorm.Open(sqlite.Open("database.db"), &gorm.Config{})
	if err != nil {
		fmt.Print("Failed " + err.Error())
	}
	// password := "admin"
	var a []models.Setting
	db.First(&a)
	fmt.Println(a)
	// hash := a.Password

	// b := bcrypt.CompareHashAndPassword([]byte(string(hash)), []byte(password))
	// fmt.Println(b)
}