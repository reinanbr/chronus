package main

import (
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/reinanbr/chronus"
)

const version = "1.0.0"

func main() {
	var (
		showVersion = flag.Bool("version", false, "Show version information")
		showHelp    = flag.Bool("help", false, "Show help information")
		sleepDur    = flag.Int("sleep", 0, "Sleep for specified milliseconds")
		showNow     = flag.Bool("now", false, "Show current time in milliseconds")
		showTime    = flag.Bool("time", false, "Execute command and show execution time")
	)

	flag.Parse()

	if *showVersion {
		fmt.Printf("chronus version %s\n", version)
		return
	}

	if *showHelp {
		printHelp()
		return
	}

	if *showNow {
		fmt.Printf("%d\n", chronus.Now())
		return
	}

	if *sleepDur > 0 {
		fmt.Printf("Sleeping for %d milliseconds...\n", *sleepDur)
		chronus.Sleep(time.Duration(*sleepDur))
		fmt.Println("Done.")
		return
	}

	if *showTime {
		if flag.NArg() == 0 {
			fmt.Fprintf(os.Stderr, "Error: no command specified for timing\n")
			os.Exit(1)
		}

		start := chronus.Now()

		// This is a simplified example - in a real CLI tool,
		// you'd want to properly execute the command with os/exec
		fmt.Printf("Timing command: %v\n", flag.Args())
		fmt.Println("(Note: actual command execution not implemented in this example)")

		end := chronus.Now()
		fmt.Printf("Execution time: %d milliseconds\n", end-start)
		return
	}

	// If no flags specified, show current time
	fmt.Printf("Current time: %d milliseconds since Unix epoch\n", chronus.Now())
}

func printHelp() {
	fmt.Printf(`chronus %s - Time utilities command line tool

Usage:
  chronus [options]

Options:
  -now          Show current time in milliseconds since Unix epoch
  -sleep N      Sleep for N milliseconds
  -time         Time the execution of a command (simplified example)
  -version      Show version information
  -help         Show this help message

Examples:
  chronus -now                    # Show current timestamp
  chronus -sleep 1000             # Sleep for 1 second
  chronus -time echo "hello"      # Time a command execution
  chronus                         # Show current time (default)

`, version)
}
