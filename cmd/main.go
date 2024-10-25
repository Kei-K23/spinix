package main

import (
	"time"

	"github.com/Kei-K23/termloader"
)

func main() {
	loader := termloader.NewLoader("spinner", 100*time.Millisecond).SetShowMessage(true).SetMessage("Loading...").SetLoaderColor("\033[32m").SetMessageColor("\033[34m")

	loader.Start()
	time.Sleep(3 * time.Second)
	loader.Stop()
}
