package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {

	slice := make([]int, 0, 3) // // Creates an empty slice with a capacity of 3

	for {
		var input string

		// Prompt the user for input
		fmt.Print("Please an integer or 'X' value to quit :")
		fmt.Scan(&input)

		// Check if the user wants to quit
		if input == "X" || input == "x" {
			break
		}

		// Convert string to integer
		num, err := strconv.Atoi(input)

		// error check
		if err != nil {
			fmt.Println("Invalid input! Please enter a valid integer or 'X' to quit.")
			continue
		}

		// adding
		slice = append(slice, num)

		// Sort the slice
		sort.Ints(slice)

		// print the sorted slice
		fmt.Println(slice)
	}

	fmt.Println()

}
