package main

import (
	"fmt"
)

// Truncate a float64 to an integer
func truncateFloat(f float64) int {
	return int(f)
}

func main() {
	var number float64

	// Prompt user for input
	fmt.Print("Please enter a floating point number: ")
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Truncate and display result
	truncated := truncateFloat(number)
	fmt.Printf("Truncated number: %d\n", truncated)
}
