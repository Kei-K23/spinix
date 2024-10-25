package termloader

import (
	"fmt"
	"sync"
	"time"
)

type Loader struct {
	style        string        // Spinner or loader
	theme        []string      // Loader animation style
	loaderColor  string        // Color for loader
	speed        time.Duration // Interval to update the loader
	message      string        // Message to show next to loader
	messageColor string        // Color for message text
	showMessage  bool          // Indicate to show message or not
	active       bool          // Status for loading indicator
	mutex        sync.Mutex    // For thread-safe
	stopCh       chan struct{} // Channel to send stop signal
}

func NewLoader(style string, speed time.Duration) *Loader {
	return &Loader{
		style:       style,
		theme:       ClassicDots,
		speed:       speed,
		active:      false,
		showMessage: true,
		stopCh:      make(chan struct{}),
	}
}

func (l *Loader) Start() {
	l.mutex.Lock()
	defer l.mutex.Unlock()

	// If already start the loader, then exist the function
	if l.active {
		return
	}

	l.active = true
	go l.animate()
}

func (l *Loader) Stop() {
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

func (l *Loader) SetMessageColor(colorCode string) *Loader {
	l.messageColor = colorCode
	return l
}

func (l *Loader) SetLoaderColor(colorCode string) *Loader {
	l.loaderColor = colorCode
	return l
}

func (l *Loader) SetMessage(message string) *Loader {
	l.message = message
	return l
}

func (l *Loader) SetShowMessage(isShow bool) *Loader {
	l.showMessage = isShow
	return l
}

func (l *Loader) SetTheme(theme []string) *Loader {
	l.theme = theme
	return l
}

func (l *Loader) animate() {
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