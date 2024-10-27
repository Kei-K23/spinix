// Package spinix provides terminal-based loading animations including spinners and progress bars.
// It supports customizable themes, colors, and speeds, allowing developers to use various styles
// to suit different terminal environments and aesthetics.
package spinix

import (
	"fmt"
	"strings"
	"sync"
	"time"
)

type Spinner struct {
	theme        []string      // Spinner animation frames
	spinnerColor string        // Color for loader
	speed        time.Duration // Interval to update the loader
	message      string        // Message to show next to loader
	messageColor string        // Color for message text
	showMessage  bool          // Indicate to show message or not
	active       bool          // Status for loading indicator
	mutex        sync.Mutex    // For thread-safe operations
	stopCh       chan struct{} // Channel to send stop signal
	callback     func()        // Optional callback to run after spinner stops
}

// NewSpinner initializes a new Spinner with default settings and theme.
func NewSpinner() *Spinner {
	return &Spinner{
		theme:        SpinnerThemes[SpinnerClassicDots], // Default theme
		speed:        100 * time.Millisecond,
		active:       false,
		spinnerColor: "\033[32m", // Green color as default
		showMessage:  true,
		stopCh:       make(chan struct{}),
	}
}

// SetCallback sets a callback function to be executed when the spinner stops.
func (s *Spinner) SetCallback(cb func()) *Spinner {
	s.callback = cb
	return s
}

// Start begins the spinner animation in a goroutine.
func (s *Spinner) Start() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if s.active {
		return
	}
	s.active = true
	go s.animate()
}

// Stop ends the spinner animation and clears the spinner line.
func (s *Spinner) Stop() {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	if !s.active {
		return
	}
	close(s.stopCh)
	s.stopCh = make(chan struct{})
	s.active = false
	fmt.Print("\r\033[K") // Clear line

	// Execute callback if set
	if s.callback != nil {
		s.callback()
	}
}

// SetMessageColor sets the color code for the spinner's message text.
func (s *Spinner) SetMessageColor(colorCode string) *Spinner {
	s.messageColor = colorCode
	return s
}

// SetSpinnerColor sets the color code for the spinner itself.
func (s *Spinner) SetSpinnerColor(colorCode string) *Spinner {
	s.spinnerColor = colorCode
	return s
}

// SetMessage sets a message to display next to the spinner.
func (s *Spinner) SetMessage(message string) *Spinner {
	s.message = message
	s.showMessage = true
	return s
}

// SetTheme selects a predefined spinner theme.
func (s *Spinner) SetTheme(theme SpinnerStyle) *Spinner {
	s.theme = SpinnerThemes[theme]
	return s
}

// SetCustomTheme allows setting a custom sequence of spinner frames.
func (s *Spinner) SetCustomTheme(frames []string) *Spinner {
	s.theme = frames
	return s
}

// SetSpeed adjusts the rotation speed of the spinner's frames.
func (s *Spinner) SetSpeed(speed time.Duration) *Spinner {
	s.speed = speed
	return s
}

// animate is an internal function that continuously updates the spinner display
// until it receives a stop signal.
func (s *Spinner) animate() {
	frameIdx := 0
	for {
		select {
		case <-s.stopCh:
			return
		default:
			fmt.Printf("\r%s%s\033[0m %s%s\033[0m", s.spinnerColor, s.theme[frameIdx], s.messageColor, s.message)
			time.Sleep(s.speed)
			frameIdx = (frameIdx + 1) % len(s.theme)
		}
	}
}

// SpinnerStyle represents a specific spinner theme.
type SpinnerStyle string

// SpinnerThemes defines several spinner animation styles.
var SpinnerThemes = map[SpinnerStyle][]string{
	SpinnerClassicDots:   {"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"},
	SpinnerLineTheme:     {"-", "\\", "|", "/"},
	SpinnerPulsatingDot:  {"⠁", "⠂", "⠄", "⠂"},
	SpinnerGrowingBlock:  {"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"},
	SpinnerRotatingArrow: {"→", "↘", "↓", "↙", "←", "↖", "↑", "↗"},
	SpinnerArcLoader:     {"◜", "◠", "◝", "◞", "◡", "◟"},
	SpinnerClock:         {"🕛", "🕐", "🕑", "🕒", "🕓", "🕔", "🕕", "🕖", "🕗", "🕘", "🕙", "🕚"},
	SpinnerCircleDots:    {"◐", "◓", "◑", "◒"},
	SpinnerBouncingBall:  {"⠁", "⠂", "⠄", "⠂"},
	SpinnerFadingSquares: {"▖", "▘", "▝", "▗"},
	SpinnerDotsFading:    {"⠁", "⠂", "⠄", "⠂", "⠁", "⠂", "⠄", "⠂"},
	SpinnerEarth:         {"🌍", "🌎", "🌏"},
	SpinnerSnake:         {"⠈", "⠐", "⠠", "⢀", "⡀", "⠄", "⠂", "⠁"},
	SpinnerTriangle:      {"◢", "◣", "◤", "◥"},
	SpinnerSpiral:        {"▖", "▘", "▝", "▗", "▘", "▝", "▖", "▗"},
	SpinnerWave:          {"▁", "▂", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃", "▂", "▁"},
	SpinnerWeather:       {"🌤️", "⛅", "🌥️", "☁️", "🌧️", "⛈️", "🌩️", "🌨️"},
	SpinnerRunningPerson: {"🏃💨", "🏃💨💨", "🏃💨💨💨", "🏃‍♂️💨", "🏃‍♂️💨💨", "🏃‍♀️💨", "🏃‍♀️💨💨"},
	SpinnerRunningCat:    {"🐱💨", "🐈💨", "🐱💨💨", "🐈💨💨"},
	SpinnerRunningDog:    {"🐕💨", "🐶💨", "🐕‍🦺💨", "🐕💨💨"},
	SpinnerCycling:       {"🚴", "🚴‍♂️", "🚴‍♀️", "🚵", "🚵‍♂️", "🚵‍♀️"},
	SpinnerCarLoading:    {"🚗💨", "🚙💨", "🚓💨", "🚕💨", "🚐💨", "🚔💨"},
	SpinnerRocket:        {"🚀", "🚀💨", "🚀💨💨", "🚀💨💨💨", "🚀🌌", "🚀🌠"},
	SpinnerOrbit:         {"🌑", "🌒", "🌓", "🌔", "🌕", "🌖", "🌗", "🌘"},
	SpinnerTrain:         {"🚆", "🚄", "🚅", "🚇", "🚊", "🚉"},
	SpinnerAirplane:      {"✈️ ", "🛫", "🛬", "✈️💨", "✈️💨💨"},
	SpinnerFireworks:     {"🎆", "🎇", "🎆🎇", "🎇🎆"},
	SpinnerPizzaDelivery: {"🍕💨", "🍔💨", "🌭💨", "🍟💨"},
	SpinnerHeartbeat:     {"💓", "💗", "💖", "💘", "💞", "💝", "💖"},
}

const (
	SpinnerClassicDots   SpinnerStyle = "classic_dots"
	SpinnerLineTheme     SpinnerStyle = "line"
	SpinnerPulsatingDot  SpinnerStyle = "pulsating_dot"
	SpinnerGrowingBlock  SpinnerStyle = "growing_block"
	SpinnerRotatingArrow SpinnerStyle = "rotating_arrow"
	SpinnerArcLoader     SpinnerStyle = "arc_loader"
	SpinnerClock         SpinnerStyle = "clock"
	SpinnerCircleDots    SpinnerStyle = "circle_dots"
	SpinnerBouncingBall  SpinnerStyle = "bouncing_ball"
	SpinnerFadingSquares SpinnerStyle = "fading_squares"
	SpinnerDotsFading    SpinnerStyle = "dots_fading"
	SpinnerEarth         SpinnerStyle = "earth"
	SpinnerSnake         SpinnerStyle = "snake"
	SpinnerTriangle      SpinnerStyle = "triangle"
	SpinnerSpiral        SpinnerStyle = "spiral"
	SpinnerWave          SpinnerStyle = "wave"
	SpinnerWeather       SpinnerStyle = "weather"
	SpinnerRunningPerson SpinnerStyle = "running_person"
	SpinnerRunningCat    SpinnerStyle = "running_cat"
	SpinnerRunningDog    SpinnerStyle = "running_dog"
	SpinnerCycling       SpinnerStyle = "cycling"
	SpinnerCarLoading    SpinnerStyle = "car_loading"
	SpinnerRocket        SpinnerStyle = "rocket"
	SpinnerOrbit         SpinnerStyle = "orbit"
	SpinnerTrain         SpinnerStyle = "train"
	SpinnerAirplane      SpinnerStyle = "airplane"
	SpinnerFireworks     SpinnerStyle = "fireworks"
	SpinnerPizzaDelivery SpinnerStyle = "pizza_delivery"
	SpinnerHeartbeat     SpinnerStyle = "heartbeat"
)

// ProgressBar represents a customizable terminal progress bar.
type ProgressBar struct {
	width          int              // Width of the progress bar
	progress       int              // Current progress percentage (0-100)
	barChar        string           // Character to display for filled progress
	emptyChar      string           // Character to display for empty progress
	leftBorder     string           // Left border of the progress bar
	rightBorder    string           // Right border of the progress bar
	color          string           // Color for progress bar
	label          string           // Optional label text
	showLabel      bool             // Flag to show or hide label
	showPercentage bool             // Flag to show or hide progress percentage
	speed          time.Duration    // Progress bar update speed
	active         bool             // Indicates if the progress bar is active
	mutex          sync.Mutex       // For thread safety
	stopCh         chan interface{} // Channel to signal stopping of the progress bar
	callback       func()           // Optional callback to run after progress bar stops
}

// NewProgressBar initializes a new ProgressBar with default settings.
func NewProgressBar() *ProgressBar {
	return &ProgressBar{
		width:          40,
		barChar:        "█",
		emptyChar:      " ",
		leftBorder:     "[",
		rightBorder:    "]",
		color:          "\033[32m", // Green color
		showPercentage: true,
		stopCh:         make(chan interface{}),
		speed:          100 * time.Millisecond,
	}
}

// SetCallback sets a callback function to be executed when the progress bar stops.
func (pb *ProgressBar) SetCallback(cb func()) *ProgressBar {
	pb.callback = cb
	return pb
}

// SetWidth sets the width of the progress bar.
func (pb *ProgressBar) SetWidth(width int) *ProgressBar {
	pb.width = width
	return pb
}

// SetLabel sets a label to display next to the progress bar.
func (pb *ProgressBar) SetLabel(label string) *ProgressBar {
	pb.label = label
	pb.showLabel = true
	return pb
}

// SetShowPercentage toggles the display of the progress percentage.
func (pb *ProgressBar) SetShowPercentage(show bool) *ProgressBar {
	pb.showPercentage = show
	return pb
}

// SetBarChar sets the character to display as filled progress.
func (pb *ProgressBar) SetBarChar(char string) *ProgressBar {
	pb.barChar = char
	return pb
}

// SetEmptyChar sets the character to display as empty progress.
func (pb *ProgressBar) SetEmptyChar(char string) *ProgressBar {
	pb.emptyChar = char
	return pb
}

// SetBorders sets the characters for the left and right borders of the progress bar.
func (pb *ProgressBar) SetBorders(left, right string) *ProgressBar {
	pb.leftBorder = left
	pb.rightBorder = right
	return pb
}

// SetSpeed sets the speed of the progress bar update interval.
func (pb *ProgressBar) SetSpeed(speed time.Duration) *ProgressBar {
	pb.speed = speed
	return pb
}

// SetColor sets the color for the progress bar.
func (pb *ProgressBar) SetColor(color string) *ProgressBar {
	pb.color = color
	return pb
}

// Start begins the progress bar animation in a separate goroutine.
func (pb *ProgressBar) Start() {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()

	if pb.active {
		return
	}
	pb.active = true
	go pb.animate()
}

// Stop ends the progress bar animation and clears the progress bar line.
func (pb *ProgressBar) Stop() {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()

	if !pb.active {
		return
	}
	close(pb.stopCh)
	pb.stopCh = make(chan interface{})
	pb.active = false
	fmt.Print("\r\033[K") // Clear line

	// Execute callback if set
	if pb.callback != nil {
		pb.callback()
	}
}

// Update sets the current progress percentage.
func (pb *ProgressBar) Update(progress int) {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()

	if progress >= 0 && progress <= 100 {
		pb.progress = progress
	}
}

// animate is an internal function that continuously updates the progress bar display.
func (pb *ProgressBar) animate() {
	for {
		select {
		case <-pb.stopCh:
			return
		default:
			pb.render()
			time.Sleep(pb.speed)
		}
	}
}

// render generates the progress bar output based on current progress.
func (pb *ProgressBar) render() {
	pb.mutex.Lock()
	defer pb.mutex.Unlock()

	fillWidth := int(float64(pb.width) * float64(pb.progress) / 100)
	emptyWidth := pb.width - fillWidth

	bar := fmt.Sprintf("\r%s%s%s%s%s\033[0m", pb.color, pb.leftBorder,
		strings.Repeat(pb.barChar, fillWidth),
		strings.Repeat(pb.emptyChar, emptyWidth), pb.rightBorder)

	if pb.showPercentage {
		bar += fmt.Sprintf(" %3d%%", pb.progress)
	}
	if pb.showLabel {
		bar += fmt.Sprintf(" %s", pb.label)
	}

	fmt.Print(bar)
}

// progressBarStyle represents the various style configurations for the progress bar.
type progressBarStyle string

const (
	PbStyleBasic      progressBarStyle = "basic"
	PbStyleClassic    progressBarStyle = "classic"
	PbStyleMinimal    progressBarStyle = "minimal"
	PbStyleBold       progressBarStyle = "bold"
	PbStyleDashed     progressBarStyle = "dashed"
	PbStyleElegant    progressBarStyle = "elegant"
	PbStyleEmoji      progressBarStyle = "emoji"
	PbStyleFuturistic progressBarStyle = "futuristic"
)

// SetStyle configures the progress bar style based on predefined styles.
func (pb *ProgressBar) SetStyle(style progressBarStyle) *ProgressBar {
	switch style {
	case PbStyleBasic:
		pb.width = 40
		pb.barChar = "="
		pb.emptyChar = "-"
		pb.leftBorder = "|"
		pb.rightBorder = "|"
		pb.color = "\033[34m" // Blue

	case PbStyleClassic:
		pb.width = 30
		pb.barChar = "#"
		pb.emptyChar = "."
		pb.leftBorder = "["
		pb.rightBorder = "]"
		pb.color = "\033[32m" // Green

	case PbStyleMinimal:
		pb.width = 20
		pb.barChar = "*"
		pb.emptyChar = " "
		pb.leftBorder = ""
		pb.rightBorder = ""
		pb.color = "\033[36m" // Cyan

	case PbStyleBold:
		pb.width = 50
		pb.barChar = "■"
		pb.emptyChar = " "
		pb.leftBorder = "❮"
		pb.rightBorder = "❯"
		pb.color = "\033[35m" // Purple

	case PbStyleDashed:
		pb.width = 45
		pb.barChar = "▮"
		pb.emptyChar = "▯"
		pb.leftBorder = "["
		pb.rightBorder = "]"
		pb.color = "\033[31m" // Red

	case PbStyleElegant:
		pb.width = 35
		pb.barChar = "▰"
		pb.emptyChar = "▱"
		pb.leftBorder = "❬"
		pb.rightBorder = "❭"
		pb.color = "\033[94m" // Light Blue

	case PbStyleEmoji:
		pb.width = 25
		pb.barChar = "🚀"
		pb.emptyChar = "✨"
		pb.leftBorder = "🚩"
		pb.rightBorder = "🎯"
		pb.color = "\033[33m" // Yellow

	case PbStyleFuturistic:
		pb.width = 40
		pb.barChar = "◉"
		pb.emptyChar = "○"
		pb.leftBorder = "⟦"
		pb.rightBorder = "⟧"
		pb.color = "\033[96m" // Cyan
	}
	return pb
}
