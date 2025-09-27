package service

import (
	"Bell/api/database"
	"context"
	"log"
	"time"

	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	Name string
	Db *gorm.DB
}

// NewApp creates a new App application struct
func NewApp(name string) *App {
	return &App{
		Name: name,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	a.Db = database.InitializeDatabase()
	err := a.Migrate()
	if err != nil {
		log.Fatal(err.Error())
	}

	weekday := time.Now().Weekday()
	wErr := a.UpdateDayActive(int8(weekday))
	if wErr != nil {
		log.Fatal(err.Error())
	}
}


