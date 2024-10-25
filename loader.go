package termloader

import (
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
