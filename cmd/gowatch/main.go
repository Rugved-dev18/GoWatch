package main

import (
	"fmt"
	"os"

	"github.com/Rugved-dev18/gowatch/internal/monitor"
)

func main() {
	file, err := os.Open("urls.txt")
	if err != nil {
		fmt.Println("Error opening urls.txt:", err)
		return
	}
	defer file.Close()

	fmt.Println("GoWatch - Website Health Monitor")
	fmt.Println("----------------------------------")

	// Temporary: we'll add proper file scanning next.
	result := monitor.CheckURL("https://google.com")

	fmt.Printf(
		"URL: %s | Status: %d | Latency: %v\n",
		result.URL,
		result.StatusCode,
		result.Latency,
	)
}
