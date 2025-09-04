package main

import (
	"fmt"

	"github.com/reinanbr/chronus"
)

func main() {
	fmt.Println("=== Chronus Library Example ===")
	// Example 1: Basic time measurement
	fmt.Println("\n1. Basic time measurement:")
	start := chronus.Now()
	fmt.Printf("Start time: %d milliseconds\n", start)

	// Sleep for 2 seconds using chronus.Sleep
	fmt.Println("Sleeping for 2 seconds...")
	chronus.Sleep(2000)

	end := chronus.Now()
	fmt.Printf("End time: %d milliseconds\n", end)

	elapsed := end - start
	fmt.Printf("Elapsed time: %d milliseconds\n", elapsed)

	// Example 2: Multiple measurements
	fmt.Println("\n2. Multiple quick measurements:")
	for i := 0; i < 3; i++ {
		timestamp := chronus.Now()
		fmt.Printf("Measurement %d: %d\n", i+1, timestamp)
		chronus.Sleep(500) // Sleep 500ms between measurements
	}

	// Example 3: Performance timing
	fmt.Println("\n3. Performance timing example:")
	perfStart := chronus.Now()

	// Simulate some work
	total := 0
	for i := 0; i < 1000000; i++ {
		total += i
	}

	perfEnd := chronus.Now()
	fmt.Printf("Calculation took: %d milliseconds\n", perfEnd-perfStart)
	fmt.Printf("Result: %d\n", total)

	fmt.Println("\n=== Example completed ===")
}
