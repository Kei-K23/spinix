# Spinix

**Spinix** 🌀 is a **Go package** that provides terminal-based loading animations, including spinners and progress bars. It supports customizable themes, colors, and speeds, allowing developers to create visually appealing loading indicators that can fit various terminal environments and aesthetics.

## Features

- **Spinners**: Multiple spinner styles with customizable speed and colors.
- **Progress Bars**: A customizable progress bar with adjustable width, characters, and styles.
- **Thread Safety**: Safe to use in concurrent environments with mutexes.
- **Customizable Themes**: Use predefined themes or create your own custom spinner themes.

## Installation

To install **Spinix**, use `go get`:

```bash
go get github.com/Kei-K23/spinix
```

## Usage

### Spinners

Here’s how to create and use a spinner:

```go
package main

import (
	"time"
	"github.com/Kei-K23/spinix"
)

func main() {
	spinner := spinix.NewSpinner().
		SetMessage("Loading...").
		SetLoaderColor("\033[34m").
		SetSpeed(100 * time.Millisecond) // Adjust speed if necessary

	spinner.Start()

	// Simulate some work
	time.Sleep(5 * time.Second)

	spinner.Stop()

    // spinix.NewSpinner() will also create spinner with default properties
    // e.g
    // spinner := spinix.NewSpinner()
	// spinner.Start()
	// time.Sleep(5 * time.Second)
	// spinner.Stop()
}

```

### Progress Bars

Creating and using a progress bar is just as straightforward:

```go
package main

import (
	"time"
	"github.com/Kei-K23/spinix"
)

func main() {
	progressBar := spinix.NewProgressBar().
		SetWidth(50).
		SetColor("\033[32m").
		SetLabel("Progress:")

	progressBar.Start()

	for i := 0; i <= 100; i++ {
		progressBar.Update(i)
		time.Sleep(50 * time.Millisecond) // Simulate work
	}

	progressBar.Stop()

    // spinix.NewProgressBar() will also create progress bar with default properties.
    // e.g
    // progressBar := spinix.NewProgressBar()
    // progressBar.Start()
	// for i := 0; i <= 100; i++ {
	// 	progressBar.Update(i)
	// 	time.Sleep(50 * time.Millisecond) // Simulate work
	// }
	// progressBar.Stop()
}
```

## Customization

### Spinners

You can customize the spinner’s message, colors, speed and theme, etc...:

```go
func main() {
	spinner := spinix.NewSpinner().
		SetMessage("Processing...").
		SetMessageColor("\033[33m").
		SetCustomTheme([]string{"▁", "▃", "▄", "▅", "▆", "▇", "█", "▇", "▆", "▅", "▄", "▃"}).
		SetSpinnerColor("\033[31m"). // Red color (you can use every color you want with that format but spinner color will not work with emoji spinner)
		SetSpeed(200 * time.Millisecond)

	spinner.Start()
	time.Sleep(2 * time.Second) // Simulate a task
	spinner.Stop()
}
```

### Progress Bars

You can customize the progress bar's message, width, color, appearance using various methods and etc...:

```go
func main() {
	progressBar := spinix.NewProgressBar().
		SetWidth(60).
		SetBarChar("█").
		SetEmptyChar("░").
		SetBorders("[", "]").
		SetShowPercentage(true)

	progressBar.Start()

	for i := 0; i <= 100; i++ {
		progressBar.Update(i)
		time.Sleep(50 * time.Millisecond) // Simulate work
	}

	progressBar.Stop()
}
```

## Use Predefined Spinner and Progress bar

### Spinners

You can also use predefined spinner themes that provided by **spinix**:
See all available predefined spinner themes below

```go
func main() {
	spinner := spinix.NewSpinner().
		SetMessage("Processing...").
		SetMessageColor("\033[33m").
		SetTheme(spinix.SpinnerRotatingArrow)

	spinner.Start()

	// Simulate some work
	time.Sleep(5 * time.Second)

	spinner.Stop()
}
```

### Progress Bars

You can use predefined progress bar style that provided by **spinix**:
See all available predefined progress bar styles below

```go
func main() {
	progressBar := spinix.NewProgressBar().
		SetStyle(spinix.PbStyleBasic).
		SetShowPercentage(true)

	progressBar.Start()

	for i := 0; i <= 100; i++ {
		progressBar.Update(i)
		time.Sleep(50 * time.Millisecond) // Simulate work
	}

	progressBar.Stop()
}
```

## Available Spinner Themes

The **Spinix** package comes with several predefined spinner themes. Here’s a list of the available styles along with their visualizations:

| Spinner Theme            | Visualization                                |
| ------------------------ | -------------------------------------------- |
| **SpinnerClassicDots**   | ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏                          |
| **SpinnerLineTheme**     | - \                                          |
| **SpinnerPulsatingDot**  | ⠁ ⠂ ⠄ ⠂                                      |
| **SpinnerGrowingBlock**  | ▁ ▃ ▄ ▅ ▆ ▇ █ ▇ ▆ ▅ ▄ ▃                      |
| **SpinnerRotatingArrow** | → ↘ ↓ ↙ ← ↖ ↑ ↗                              |
| **SpinnerArcLoader**     | ◜ ◠ ◝ ◞ ◡ ◟                                  |
| **SpinnerClock**         | 🕛 🕐 🕑 🕒 🕓 🕔 🕕 🕖 🕗 🕘 🕙 🕚          |
| **SpinnerCircleDots**    | ◐ ◓ ◑ ◒                                      |
| **SpinnerBouncingBall**  | ⠁ ⠂ ⠄ ⠂                                      |
| **SpinnerFadingSquares** | ▖ ▘ ▝ ▗                                      |
| **SpinnerDotsFading**    | ⠁ ⠂ ⠄ ⠂ ⠁ ⠂ ⠄ ⠂                              |
| **SpinnerEarth**         | 🌍 🌎 🌏                                     |
| **SpinnerSnake**         | ⠈ ⠐ ⠠ ⢀ ⡀ ⠄ ⠂ ⠁                              |
| **SpinnerTriangle**      | ◢ ◣ ◤ ◥                                      |
| **SpinnerSpiral**        | ▖ ▘ ▝ ▗ ▘ ▝ ▖ ▗                              |
| **SpinnerWave**          | ▁ ▂ ▃ ▄ ▅ ▆ ▇ █ ▇ ▆ ▅ ▄ ▃ ▂ ▁                |
| **SpinnerWeather**       | 🌤️ ⛅ 🌥️ ☁️ 🌧️ ⛈️ 🌩️ 🌨️                      |
| **SpinnerRunningPerson** | 🏃💨 🏃💨💨 🏃💨💨💨 🏃‍♂️💨 🏃‍♂️💨💨 🏃‍♀️💨 🏃‍♀️💨💨 |
| **SpinnerRunningCat**    | 🐱💨 🐈💨 🐱💨💨 🐈💨💨                      |
| **SpinnerRunningDog**    | 🐕💨 🐶💨 🐕‍🦺💨 🐕💨💨                        |
| **SpinnerCycling**       | 🚴 🚴‍♂️ 🚴‍♀️ 🚵 🚵‍♂️ 🚵‍♀️                            |
| **SpinnerCarLoading**    | 🚗💨 🚙💨 🚓💨 🚕💨 🚐💨 🚔💨                |
| **SpinnerRocket**        | 🚀 🚀💨 🚀💨💨 🚀💨💨💨 🚀🌌 🚀🌠            |
| **SpinnerOrbit**         | 🌑 🌒 🌓 🌔 🌕 🌖 🌗 🌘                      |
| **SpinnerTrain**         | 🚆 🚄 🚅 🚇 🚊 🚉                            |
| **SpinnerAirplane**      | ✈️ 🛫 🛬 ✈️💨 ✈️💨💨                         |
| **SpinnerFireworks**     | 🎆 🎇 🎆🎇 🎇🎆                              |
| **SpinnerPizzaDelivery** | 🍕💨 🍔💨 🌭💨 🍟💨                          |
| **SpinnerHeartbeat**     | 💓 💗 💖 💘 💞 💝 💖                         |

## Available Progress Bar Styles

The ProgressBar can be styled using various predefined styles. Here’s a list of available styles:

| Progress Bar Style    | Description                                                |
| --------------------- | ---------------------------------------------------------- |
| **PbStyleBasic**      | ===========================------------- 69%               |
| **PbStyleClassic**    | [############################..] 95%                       |
| **PbStyleMinimal**    | **\*\***\*\*\***\*\*** 79%                                 |
| **PbStyleBold**       | ❮■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■■ ❯ 93%      |
| **PbStyleDashed**     | [▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▮▯▯▯▯▯▯▯▯▯▯▯▯▯▯▯▯▯▯] 61%        |
| **PbStyleElegant**    | ❬▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▰▱▱▱▱▱▱▱▱❭ 79                   |
| **PbStyleEmoji**      | 🚩🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀🚀✨✨✨✨✨✨✨✨🎯 71% |
| **PbStyleFuturistic** | ⟦◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉◉○○○○○○○○○○○○○○○○○⟧ 59%             |

## Example

Here’s a complete example that demonstrates both the spinner and the progress bar together:

```go
package main

import (
	"time"
	"github.com/Kei-K23/spinix"
)

func main() {
	spinner := spinix.NewSpinner()
	progressBar := spinix.NewProgressBar()

	spinner.SetMessage("Loading Data...")
	spinner.Start()
	progressBar.SetWidth(50)
	progressBar.SetColor("\033[36m") // Cyan
	progressBar.Start()

	for i := 0; i <= 100; i++ {
		progressBar.Update(i)
		time.Sleep(50 * time.Millisecond) // Simulate work
	}

	progressBar.Stop()
	spinner.Stop()
}
```

## Contributing

Contributions are welcome! Please feel free to open issues, submit pull requests, or provide feedback.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.
