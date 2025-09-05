package chronus

import "time"

// Now returns the current time in milliseconds since the Unix epoch (January 1, 1970 UTC).
//
// This function provides a convenient way to get high-precision timestamps for timing
// operations, logging, or any application that requires millisecond-level time tracking.
// The returned value is always positive and increases monotonically.
//
// Returns:
//   - int64: The current Unix timestamp in milliseconds
//
// Example:
//
//	timestamp := chronus.Now()
//	fmt.Printf("Current time: %d milliseconds\n", timestamp)
//
// For timing operations:
//
//	start := chronus.Now()
//	// ... some operation
//	duration := chronus.Now() - start
//	fmt.Printf("Operation took %d ms\n", duration)
func Now() int64 {
	return time.Now().UnixMilli()
}
