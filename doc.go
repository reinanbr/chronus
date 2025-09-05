// Package chronus provides time utilities with millisecond precision for Go applications.
//
// Chronus offers simple and efficient functions for working with time in milliseconds,
// including getting current timestamps and sleeping with millisecond precision.
//
// # Features
//
//   - Get current time in milliseconds since Unix epoch
//   - Sleep functionality with millisecond precision
//   - Simple and intuitive API
//   - Zero dependencies beyond Go standard library
//   - High performance with minimal overhead
//
// # Basic Usage
//
// Getting the current time in milliseconds:
//
//	timestamp := chronus.Now()
//	fmt.Printf("Current time: %d milliseconds\n", timestamp)
//
// Sleeping for a specific duration:
//
//	chronus.Sleep(1000) // Sleep for 1 second
//	chronus.Sleep(500)  // Sleep for 500 milliseconds
//
// # Timing Operations
//
// You can use chronus to measure execution time:
//
//	start := chronus.Now()
//	// ... perform some operation
//	elapsed := chronus.Now() - start
//	fmt.Printf("Operation took %d milliseconds\n", elapsed)
//
// # Performance
//
// Chronus functions are designed to be lightweight and fast, suitable for
// high-frequency timing operations without significant overhead.
package chronus
