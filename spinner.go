package termloader

import (
	"fmt"
	"sync"
	"time"
)

type Spinner struct {
	theme        []string      // Spinner animation style
	loaderColor  string        // Color for loader
	speed        time.Duration // Interval to update the loader
	message      string        // Message to show next to loader
	messageColor string        // Color for message text
	showMessage  bool          // Indicate to show message or not
	active       bool          // Status for loading indicator
	mutex        sync.Mutex    // For thread-safe
	stopCh       chan struct{} // Channel to send stop signal
}

func NewSpinner() *Spinner {
	return &Spinner{
		theme:       ClassicDots,
		speed:       100 * time.Millisecond,
		active:      false,
		loaderColor: "\033[32m", // Green color as default,
		showMessage: true,
		stopCh:      make(chan struct{}),
	}
}

func (l *Spinner) Start() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// If already start the loader, then exist the function
	if l.active {
		return
	}

	l.active = true
	go l.animate()
}

func (l *Spinner) Stop() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	if !l.active {
		return
	}

	close(l.stopCh)                // Send close signal
	l.stopCh = make(chan struct{}) // Reset channel
	l.active = false
	fmt.Print("\r\033[K") // Clear line after stopping
}

func (l *Spinner) SetMessageColor(colorCode string) *Spinner {
	l.messageColor = colorCode
	return l
}

func (l *Spinner) SetLoaderColor(colorCode string) *Spinner {
	l.loaderColor = colorCode
	return l
}

func (l *Spinner) SetMessage(message string) *Spinner {
	l.message = message
	return l
}

func (l *Spinner) SetShowMessage(isShow bool) *Spinner {
	l.showMessage = isShow
	return l
}

func (l *Spinner) SetTheme(theme []string) *Spinner {
	l.theme = theme
	return l
}

func (l *Spinner) SetSpeed(speed time.Duration) *Spinner {
	l.speed = speed
	return l
}

func (l *Spinner) animate() {
	frameIdx := 0
	for {
		select {
		case <-l.stopCh:
			return
		default:
			fmt.Printf("\r%s%s\033[0m %s%s\033[0m", l.loaderColor, l.theme[frameIdx], l.messageColor, l.message)
			time.Sleep(l.speed)
			frameIdx = (frameIdx + 1) % len(l.theme)
		}
	}
}
