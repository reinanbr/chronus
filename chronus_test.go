package chronus

import (
	"testing"
	"time"
)

func TestNow(t *testing.T) {
	// Test that Now() returns a valid timestamp
	now := Now()
	if now <= 0 {
		t.Errorf("Expected positive timestamp, got %d", now)
	}

	// Test that Now() returns increasing values
	now1 := Now()
	time.Sleep(1 * time.Millisecond)
	now2 := Now()

	if now2 <= now1 {
		t.Errorf("Expected now2 (%d) to be greater than now1 (%d)", now2, now1)
	}

	// Test that the timestamp is reasonable (within the last 100 years and next 100 years)
	// Unix epoch in milliseconds: 1970-01-01 = 0
	// Year 1920 ≈ -1577923200000, Year 2120 ≈ 4732713600000
	minTime := int64(-1577923200000)
	maxTime := int64(4732713600000)

	if now < minTime || now > maxTime {
		t.Errorf("Timestamp %d seems unreasonable (not between %d and %d)", now, minTime, maxTime)
	}
}

func TestSleep(t *testing.T) {
	// Test basic sleep functionality
	start := time.Now()
	Sleep(100) // Sleep for 100 milliseconds
	elapsed := time.Since(start)

	// Allow some tolerance for timing (±50ms)
	expectedMin := 50 * time.Millisecond
	expectedMax := 150 * time.Millisecond

	if elapsed < expectedMin || elapsed > expectedMax {
		t.Errorf("Expected sleep duration between %v and %v, got %v", expectedMin, expectedMax, elapsed)
	}
}

func TestSleepZero(t *testing.T) {
	// Test that zero duration sleep works
	start := time.Now()
	Sleep(0)
	elapsed := time.Since(start)

	// Should be very quick (less than 10ms)
	if elapsed > 10*time.Millisecond {
		t.Errorf("Expected zero sleep to be very quick, took %v", elapsed)
	}
}

func BenchmarkNow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Now()
	}
}

func BenchmarkSleep(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sleep(1) // Sleep for 1 millisecond
	}
}

// Example demonstrating the usage of Now function
func ExampleNow() {
	timestamp := Now()
	_ = timestamp // timestamp contains current time in milliseconds
}

// Example demonstrating the usage of Sleep function
func ExampleSleep() {
	Sleep(1000) // Sleep for 1000 milliseconds (1 second)
}
