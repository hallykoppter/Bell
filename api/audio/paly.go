package audio

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/faiface/beep"
	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

const DefaultSampleRate = beep.SampleRate(44100)

var initialized = false

func InitSpeaker() {
	if !initialized {
		bufferSize := DefaultSampleRate.N(time.Second / 10) 
		err := speaker.Init(DefaultSampleRate, bufferSize)
		if err != nil {
			fmt.Printf("Gagal inisialisasi speaker: %v", err)
		}
		initialized = true
		fmt.Println("Speaker driver berhasil diinisialisasi.")
	}
}

func PlayBellSound(soundFile string) error {
	if !initialized {
		return fmt.Errorf("speaker belum diinisialisasi")
	}
	

	f, err := os.Open(filepath.Join("sound", soundFile))
	if err != nil {
		 return fmt.Errorf("gagal membuka file audio %s: %w", soundFile, err)
	}
	defer f.Close()

	streamer, format, err := mp3.Decode(f)
	if err != nil {
		 return fmt.Errorf("gagal decode file audio: %w", err)
	}
	defer streamer.Close()

	var s beep.Streamer = streamer
	if format.SampleRate != DefaultSampleRate {
		fmt.Println("Resampling audio")
		s = beep.Resample(3, format.SampleRate, DefaultSampleRate, streamer)
	}

	// speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))

	var done = make(chan bool)
	speaker.Play(beep.Seq(s, beep.Callback(func() {
		done <- true
	})))

	<-done

	return nil
}

func StopBellSound() {
	if !initialized {
		fmt.Println("Speaker belum diinisialisasi, tidak ada audio yang dihentikan")
		return
	}
	speaker.Clear()

	fmt.Println("Pemutaran audio telah dihentikan.")
}