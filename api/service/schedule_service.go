package service

import (
	"Bell/api/audio"
	"Bell/api/models"
	"Bell/api/repository"
	"fmt"
	"time"
)

func (a *App) InsertSchedule(Waktu string, BellID, DayID int) int {
	schedule := models.Schedule{
		Waktu: Waktu,
		BellID: uint(BellID),
		DayID: uint(DayID),
	}
	res := repository.QueryInsertSchedule(a.Db, &schedule)
	return res
}

func (a *App) GetScheduleByDay(ID int) ([]models.Schedule, error) {
	res, err := repository.QueryScheduleByDay(a.Db, ID)
	return res, err
}

func (a *App) DeleteScheduleByID(ID int) error {
	err := repository.QueryScheduleDelete(a.Db, ID)
	return err
}

func (a *App) StopBell() {
	audio.StopBellSound()
}

// Mengubah hari menjadi ID dengan asumsi 1=Senin dan 7=minggu
func GoDayToDBID(wd time.Weekday) uint {
	if wd == time.Sunday {
		wd = 7
	}
	return uint(wd)
}


// Fungsi Bel Otomatis di latar belakang
func (a *App) RunScheduler(){

	// mengambil waktu per detik
	ticker := time.NewTicker(time.Second * 1)
	defer ticker.Stop()

	// set untuk melacak bel yang sudah dimainkan hari ini
	playedBells := make(map[uint]bool)

	// Loop goroutine
	for range ticker.C {
		now := time.Now()

		// cek apakah hari sudah berganti atau belum dan reset playedBell
		if now.Hour() == 0 && now.Minute() == 0 && now.Second() < 5 {
			playedBells = make(map[uint]bool)
		}

		setting, err := repository.QuerySetting(a.Db)
		if err != nil {
			fmt.Printf("error fetching setting %v", err)
		}

		currentDayID := setting.DayActive
		curentTimeStr := now.Format("15:04:00")

		// 1 Ambil semua jadwal hari ini
		// query lah semua jadwal di hari ini
		schedules, err := repository.QueryScheduleByDay(a.Db, int(currentDayID))
		if err != nil {
			fmt.Printf("error fetching schedule %v", err)
			continue
		}

		if len(schedules) == 0 {
            fmt.Printf("ℹ️ Jadwal bel hari ini (DayID: %d) belum disiapkan. Melewati cek.", currentDayID)
            continue // Lanjut ke tick berikutnya
        }

		// 2 Cek setiap jadwalnya
		for _, schedule := range schedules {
			if schedule.Waktu == curentTimeStr && !playedBells[schedule.ID] {
				
				// 3 putar audio
				err := audio.PlayBellSound(schedule.Bell.SoundFile)
				if err != nil {
					fmt.Printf("Gagal memutar audion %v", err)
				} else {
					fmt.Println("Berhasil memutar audio")
					playedBells[schedule.ID] = true
				}
			}
		}
	}
}