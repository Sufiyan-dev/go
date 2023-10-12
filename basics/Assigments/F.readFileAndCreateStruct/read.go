package main

import (
	"fmt"
	"os"
	"strings"
)

// Error checker
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	fileName := ""
	type person struct {
		fname string
		lname string
	}

	persons := make([]person, 0)

	fmt.Print("Enter file name with directory path ")
	fmt.Scan(&fileName)

	content, err := os.ReadFile(fileName)

	check(err)

	// fmt.Println(string(content))

	lines := strings.Split(string(content), "\n")

	// fmt.Println(lines)

	for _, line := range lines {

		parts := strings.Fields(line)
		fmt.Println(parts)
		if len(parts) == 2 {
			persons = append(persons, person{fname: parts[0], lname: parts[1]})
		} else {
			fmt.Print("Invalid names, expected 2 names in each lines")
			return
		}

	}

	for i, person := range persons {
		fmt.Println(i, "First name :", person.fname, ", Last name :", person.lname)
	}
}
