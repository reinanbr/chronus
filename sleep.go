package chronus

import "time"

// Sleep pauses the current goroutine for at least the specified duration in milliseconds.
//
// This function provides a convenient wrapper around time.Sleep with millisecond precision,
// making it easier to work with millisecond-based timing without manual conversion.
// The function blocks the current goroutine and does not return until the specified
// duration has elapsed.
//
// Parameters:
//   - duration: The number of milliseconds to sleep (as time.Duration)
//
// Example:
//
//	chronus.Sleep(1000) // Sleep for 1 second (1000 milliseconds)
//	chronus.Sleep(500)  // Sleep for half a second
//
// Common usage patterns:
//
//	// Short delay
//	chronus.Sleep(100)
//
//	// Rate limiting
//	for i := 0; i < 10; i++ {
//		// do work
//		chronus.Sleep(200) // 200ms between iterations
//	}
func Sleep(duration time.Duration) {
	time.Sleep(duration * time.Millisecond)
}
