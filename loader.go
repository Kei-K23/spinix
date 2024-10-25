package termloader

import (
	"fmt"
	"sync"
	"time"
)

type Loader struct {
	Style       string        // Spinner or loader
	Theme       []string      // Loader animation style
	Speed       time.Duration // Interval to update the loader
	Message     string        // Message to show next to loader
	ShowMessage bool          // Indicate to show message or not
	Active      bool          // Status for loading indicator
	Mutex       sync.Mutex    // For thread-safe
	StopCh      chan struct{} // Channel to send stop signal
}

func NewLoader(style string, theme []string, speed time.Duration) *Loader {
	return &Loader{
		Style:  style,
		Theme:  theme,
		Speed:  speed,
		Active: false,
		StopCh: make(chan struct{}),
	}
}

func (l *Loader) Start() {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	// If already start the loader, then exist the function
	if l.Active {
		return
	}

	l.Active = true
	go l.animate()
}

func (l *Loader) Stop() {
	l.Mutex.Lock()
	defer l.Mutex.Unlock()

	if !l.Active {
		return
	}

	close(l.StopCh)                // Send close signal
	l.StopCh = make(chan struct{}) // Reset channel
	l.Active = false
	fmt.Print("\r\033[K") // Clear line after stopping
}

func (l *Loader) SetColor(colorCode string) {
	l.Message = colorCode + l.Message + "\033[0m"
}

func (l *Loader) SetTheme(theme []string) {
	l.Theme = theme
}

func (l *Loader) animate() {
	frameIdx := 0
	for {
		select {
		case <-l.StopCh:
			return
		default:
			fmt.Printf("\r%s %s\n", l.Theme[frameIdx], l.Message)
			time.Sleep(l.Speed)
			frameIdx = (frameIdx + 1) % len(l.Theme)
		}
	}
}
