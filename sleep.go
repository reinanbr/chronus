package chronus

import "time"

// Sleep pauses the current goroutine for at least the specified duration in milliseconds.
// This function provides a convenient wrapper around time.Sleep with millisecond precision,
// making it easier to work with millisecond-based timing without manual conversion.
//
// Parameters:
//   - duration: The number of milliseconds to sleep
//
// Example:
//
//	chronus.Sleep(1000) // Sleep for 1 second (1000 milliseconds)
//	chronus.Sleep(500)  // Sleep for half a second
func Sleep(duration time.Duration) {
	time.Sleep(duration * time.Millisecond)
}
