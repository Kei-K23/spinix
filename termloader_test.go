// Package termloader provides terminal-based loading animations including spinners and progress bars.
package termloader

import (
	"testing"
	"time"
)

// TestSpinner tests the basic functionalities of the Spinner.
func TestSpinner(t *testing.T) {
	spinner := NewSpinner()

	// Test default values
	if spinner.speed != 100*time.Millisecond {
		t.Errorf("Expected speed 100ms, got %v", spinner.speed)
	}
	if spinner.loaderColor != "\033[32m" {
		t.Errorf("Expected loaderColor green, got %s", spinner.loaderColor)
	}
	if spinner.active {
		t.Error("Expected spinner to be inactive initially")
	}

	// Test starting the spinner
	spinner.Start()
	time.Sleep(150 * time.Millisecond) // Allow some time for spinner animation
	if !spinner.active {
		t.Error("Expected spinner to be active after starting")
	}

	// Stop the spinner
	spinner.Stop()
	if spinner.active {
		t.Error("Expected spinner to be inactive after stopping")
	}
}

// TestProgressBar tests the basic functionalities of the ProgressBar.
func TestProgressBar(t *testing.T) {
	pb := NewProgressBar()

	// Test default values
	if pb.width != 40 {
		t.Errorf("Expected width 40, got %d", pb.width)
	}
	if pb.progress != 0 {
		t.Errorf("Expected initial progress 0, got %d", pb.progress)
	}
	if pb.barChar != "█" {
		t.Errorf("Expected barChar '█', got %s", pb.barChar)
	}
	if pb.active {
		t.Error("Expected progress bar to be inactive initially")
	}

	// Start the progress bar
	pb.Start()
	time.Sleep(150 * time.Millisecond) // Allow some time for progress bar animation
	if !pb.active {
		t.Error("Expected progress bar to be active after starting")
	}

	// Test updating progress
	pb.Update(50)
	if pb.progress != 50 {
		t.Errorf("Expected progress to be 50, got %d", pb.progress)
	}

	// Test stopping the progress bar
	pb.Stop()
	if pb.active {
		t.Error("Expected progress bar to be inactive after stopping")
	}
}
