package service

import (
	"Bell/api/models"
	"fmt"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
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

func TestGetDay(t *testing.T) {
	weekday := time.Now().Weekday()
	fmt.Println(weekday)
	fmt.Println(int(weekday))
}

func TestRingBell(t *testing.T) {
	soundFile := "pembuka.mp3"
	f, err := os.Open(filepath.Join("../../sound", soundFile))
	if err != nil {
		fmt.Println("gagal decode file: ",err)
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		fmt.Println(err)
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	done := make(chan bool)
	speaker.Play(beep.Seq(streamer, beep.Callback(func () {
		done <- true
	})))

	<- done
	fmt.Printf("bel sudah berbunyi: %s", soundFile)
}