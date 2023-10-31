package main

import (
	"fmt"
	"strings"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

func (a animal) Eat() string {
	return a.food
}

func (a animal) Move() string {
	return a.locomotion
}

func (a animal) Speak() string {
	return a.noise
}

func main() {
	animals := map[string]animal{
		"cow":   {"grass", "walk", "moo"},
		"bird":  {"worms", "fly", "peep"},
		"snake": {"mice", "slither", "hsss"},
	}

	for {
		fmt.Println("Enter animal name or 'X' to exit :")

		var animalName, infoType string

		fmt.Print("> ")
		fmt.Scan(&animalName)

		if animalName == "X" {
			break
		}

		animal, ok := animals[strings.ToLower(animalName)]

		if !ok {
			fmt.Println("Invalid animal name")
			continue
		} else {
			for {

				fmt.Println("Enter method or 'X' to exit method:")
				fmt.Scan(&infoType)

				if infoType == "X" {
					break
				}

				switch strings.ToLower(infoType) {

				case "eat":
					fmt.Println(animal.Eat())

				case "locomotive":
					fmt.Println(animal.Move())

				case "noise":
					fmt.Println(animal.food)

				default:
					fmt.Println("Invalid method")
				}

			}
		}

	}
}
