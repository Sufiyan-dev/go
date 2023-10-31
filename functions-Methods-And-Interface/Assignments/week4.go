package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

// Cow
type Cow struct {
	name string
}

func (c Cow) Eat() {
	fmt.Println("Grass")
}
func (c Cow) Move() {
	fmt.Println("Walk")
}
func (c Cow) Speak() {
	fmt.Println("Moo")
}

// Snake
type Snake struct {
	name string
}

func (s Snake) Eat() {
	fmt.Println("Mice")
}
func (s Snake) Move() {
	fmt.Println("Slither")
}
func (s Snake) Speak() {
	fmt.Println("Hsss")
}

// Bird
type Bird struct {
	name string
}

func (b Bird) Eat() {
	fmt.Println("Worms")
}
func (b Bird) Move() {
	fmt.Println("Fly")
}
func (b Bird) Speak() {
	fmt.Println("Peep")
}

func main() {

	m := make(map[string]Animal)
	cowName := Cow{}
	birdName := Bird{}
	snakeName := Snake{}

	fmt.Println("Type 'newanimal' followed by animal name and its animal type to add new animal")
	fmt.Println("Type 'query' followed by animal name and its property name to see the property")
	fmt.Println("Type Exit to exit")

	for {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			line := scanner.Text()
			if strings.ToLower(line) == "exit" {
				break
			}
			keys := strings.Split(line, " ")
			if len(keys) != 3 {
				fmt.Println("Invalid entry")
				continue
			}
			if strings.ToLower(keys[0]) == "newanimal" {
				if strings.ToLower(keys[2]) == "cow" {
					m[strings.ToLower(keys[1])] = cowName
				} else if strings.ToLower(keys[2]) == "bird" {
					m[strings.ToLower(keys[1])] = birdName
				} else if strings.ToLower(keys[2]) == "snake" {
					m[strings.ToLower(keys[1])] = snakeName
				} else {
					fmt.Println("Enter either cow, bird or snake as animal type")
					continue
				}
				fmt.Println("Created it!")

			} else if strings.ToLower(keys[0]) == "query" {
				animal, yesNo := m[strings.ToLower(keys[1])]
				if !yesNo {
					fmt.Println("Animal name not present.")
					continue
				}
				if strings.ToLower(keys[2]) == "eat" {
					animal.Eat()
				} else if strings.ToLower(keys[2]) == "move" {
					animal.Move()
				} else if strings.ToLower(keys[2]) == "speak" {
					animal.Speak()
				} else {
					fmt.Println("Enter 'eat', 'move' or 'speak'")
				}

			} else {
				fmt.Println("Enter either 'newanimal' or 'query'!")
				continue
			}
		}
	}
}
