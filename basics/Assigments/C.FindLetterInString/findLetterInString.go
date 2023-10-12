package main

import (
	"fmt"
	"strings"
)

func FindLetter(word string, letterToFind string) bool {
	exist := strings.Contains(word, letterToFind)
	return exist
}

func main() {

	var word string

	fmt.Print("Please enter a string: ")
	fmt.Scan(&word)

	word = strings.ToLower(word)

	if strings.HasPrefix(word, "i") && strings.HasSuffix(word, "n") && strings.Contains(word, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
