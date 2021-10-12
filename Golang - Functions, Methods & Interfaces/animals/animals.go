package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food, locomotion, speech string
}

func (a Animal) Eat() string {
	return a.food
}
func (a Animal) Move() string {
	return a.locomotion
}
func (a Animal) Speak() string {
	return a.speech
}

func main() {

	var animal string
	var activity string

	cow := new(Animal)
	cow.food = "Grass"
	cow.locomotion = "Walk"
	cow.speech = "Moo"

	bird := new(Animal)
	bird.food = "Worms"
	bird.locomotion = "Fly"
	bird.speech = "Peep"

	snake := new(Animal)
	snake.food = "Mice"
	snake.locomotion = "Slither"
	snake.speech = "Hsss"

	for true {
		fmt.Printf("Insert your animal (cow, bird, snake) and it's activity (eat, move, speak):\n>")
		fmt.Scan(&animal, &activity)

		switch {
		case strings.ToLower(animal) == "cow" && strings.ToLower(activity) == "eat":
			fmt.Println(cow.Eat())

		case strings.ToLower(animal) == "cow" && strings.ToLower(activity) == "move":
			fmt.Println(cow.Move())

		case strings.ToLower(animal) == "cow" && strings.ToLower(activity) == "speak":
			fmt.Println(cow.Speak())

		case strings.ToLower(animal) == "bird" && strings.ToLower(activity) == "eat":
			fmt.Println(bird.Eat())

		case strings.ToLower(animal) == "bird" && strings.ToLower(activity) == "move":
			fmt.Println(bird.Move())

		case strings.ToLower(animal) == "bird" && strings.ToLower(activity) == "speak":
			fmt.Println(bird.Speak())

		case strings.ToLower(animal) == "snake" && strings.ToLower(activity) == "eat":
			fmt.Println(snake.Eat())

		case strings.ToLower(animal) == "snake" && strings.ToLower(activity) == "move":
			fmt.Println(snake.Move())

		case strings.ToLower(animal) == "snake" && strings.ToLower(activity) == "speak":
			fmt.Println(snake.Speak())

		default:
			fmt.Println("That animal is not valid")
		}
	}
}
