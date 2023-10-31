package main

import (
	"fmt"
)

type Animal interface {
	Eat() string
	Move() string
	Speak() string
}

type Cow struct {
	eat   string
	move  string
	speak string
}

func (c Cow) Eat() string {
	return c.eat
}

func (c Cow) Move() string {
	return c.move
}

func (c Cow) Speak() string {
	return c.speak
}

type Bird struct {
	eat   string
	move  string
	speak string
}

func (b Bird) Eat() string {
	return b.eat
}

func (b Bird) Move() string {
	return b.move
}

func (b Bird) Speak() string {
	return b.speak
}

type Snake struct {
	eat   string
	move  string
	speak string
}

func (s Snake) Eat() string {
	return s.eat
}

func (s Snake) Move() string {
	return s.move
}

func (s Snake) Speak() string {
	return s.speak
}

func main() {
	animalMap := make(map[string]Animal)

	for {
		fmt.Println("> ")
		var command, name, action string
		fmt.Scan(&command, &name, &action)

		switch command {
		case "newanimal":
			animalType := action
			switch animalType {
			case "cow":
				animalMap[name] = Cow{"grass", "walk", "moo"}
			case "bird":
				animalMap[name] = Bird{"worms", "fly", "peep"}
			case "snake":
				animalMap[name] = Snake{"mice", "slither", "hsss"}
			default:
				fmt.Print("Invalid action")
			}
			fmt.Println("Created it!")
		case "query":
			animal, exists := animalMap[name]
			if !exists {
				fmt.Println("Animal not found")
				continue
			}

			switch action {
			case "eat":
				fmt.Print(animal.Eat())
			case "move":
				fmt.Print(animal.Move())
			case "speak":
				fmt.Print(animal.Speak())
			}
		default:
			fmt.Println("Invalid command")
		}
	}
}
