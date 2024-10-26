package main

import (
	"time"

	"github.com/Kei-K23/termloader"
)

func main() {
	progress := termloader.NewProgressBar().
		SetWidth(50).
		SetStyle(termloader.PbStyleEmoji).
		SetColor("\033[32m"). // Blue color
		SetShowPercentage(true).
		SetSpeed(100 * time.Millisecond)

	progress.Start()

	// Simulate some work
	for i := 0; i <= 100; i += 5 {
		progress.Update(i)
		time.Sleep(200 * time.Millisecond)
	}
	progress.Stop()

	// spinner := termloader.NewSpinner().SetCustomTheme([]string{"🔅", "🔆", "🔅", "🔆"})
	// spinner.Start()
	// time.Sleep(2 * time.Second)
	// spinner.Stop()
}
