package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	fmt.Println("--- Cloud Build Sandbox Demo ---")
	fmt.Printf("Current time: %s\n", time.Now().Format(time.RFC3339))

	hostname, err := os.Hostname()
	if err != nil {
		hostname = "unknown"
	}
	fmt.Printf("Hostname: %s\n", hostname)
	fmt.Println("\nBuild test successful!")
}
