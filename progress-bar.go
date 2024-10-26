package termloader

import (
	"sync"
	"time"
)

type ProgressBar struct {
	width          int              // Width of the progress bar
	progress       int              // Current progress percentage (based on 100%)
	barChar        string           // Character to display a filled part
	emptyChar      string           // Character to display a empty part
	leftBorder     string           // Character to use as progress left border
	rightBorder    string           // Character to use as progress right border
	color          string           // Color for progress bar
	label          string           // Optional label text to show
	showLabel      bool             // Flag to show or hide label
	showPercentage bool             // Flag to show or hide progress percentage
	speed          time.Duration    // Progress bar speed
	active         bool             // Indicator for progress bar is running/active progress or finished
	mutex          sync.Mutex       // For thread safe
	stopCh         chan interface{} // Signal to send to stop the go routine and stop the progress bar
}

func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		width:          40,
		barChar:        "█",
		emptyChar:      " ",
		leftBorder:     "[",
		rightBorder:    "]",
		color:          "\033[32m", // Green color as default
		showPercentage: true,
		stopCh:         make(chan interface{}),
		speed:          100 * time.Millisecond,
	}
}

func (pb *ProgressBar) Update(progress int) {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()

	// Progress need to be within 100%
	if pb.progress >= 0 && pb.progress <= 100 {
		pb.progress = progress
	}
}
