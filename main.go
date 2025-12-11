package main

import (
	"fmt"
	"os"
	"time"
)

// feat

// hostnameGetter is a function type that returns hostname and error
type hostnameGetter func() (string, error)

var dagetter hostnameGetter = os.Hostname

func getHostname() string {
	return getHostnameWithGetter(dagetter)
}

func getHostnameWithGetter(getter hostnameGetter) string {
	hostname, err := getter()
	if err != nil {
		return "unknown"
	}
	return hostname
}

func formatTimestamp(t time.Time) string {
	return t.Format(time.RFC3339)
}

func generateMessage() string {
	return "--- Cloud Build Sandbox Demo ---"
}

func main() {
	fmt.Println(generateMessage())
	fmt.Printf("Current time: %s\n", formatTimestamp(time.Now()))
	fmt.Printf("Hostname: %s\n", getHostname())
	fmt.Println("\nBuild test successful!")
}
