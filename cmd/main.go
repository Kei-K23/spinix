package main

import (
	"time"

	"github.com/Kei-K23/termloader"
)

func main() {
	spinner := termloader.NewSpinner()

	spinner.Start()
	time.Sleep(3 * time.Second)
	spinner.Stop()
}
