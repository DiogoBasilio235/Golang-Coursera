package main

import (
	"fmt"
)

type Animal interface {
	Eat()
	Move()
	Speak()
	getName() string
}

type Cow struct{ name string }
type Bird struct{ name string }
type Snake struct{ name string }

func (c Cow) Eat()            { fmt.Println("Grass") }
func (c Cow) Move()           { fmt.Println("Walk") }
func (c Cow) Speak()          { fmt.Println("Mooo") }
func (c Cow) getName() string { return c.name }

func (b Bird) Eat()            { fmt.Println("Worms") }
func (b Bird) Move()           { fmt.Println("Fly") }
func (b Bird) Speak()          { fmt.Println("Peep") }
func (b Bird) getName() string { return b.name }

func (s Snake) Eat()            { fmt.Println("Mice") }
func (s Snake) Move()           { fmt.Println("Slither") }
func (s Snake) Speak()          { fmt.Println("Hsss") }
func (s Snake) getName() string { return s.name }

func main() {
	var inRequest string
	var inNameOrAnimal string
	var inTypeOrAct string

	var animalSlice = []Animal{}

	fmt.Printf("Insert your request(newanimal or query)\n")
	fmt.Printf("If 'newanimal', insert your animal name and it's type (cow, bird, snake) \n")
	fmt.Printf("If 'query', insert animal name and the activity you want to know\n")

	for true {
		fmt.Printf(">")
		fmt.Scan(&inRequest, &inNameOrAnimal, &inTypeOrAct)

		switch inRequest {

		case "newanimal":
			if inTypeOrAct == "cow" {
				animalSlice = append(animalSlice, Cow{name: inNameOrAnimal})
				fmt.Println("New type of land animal created")

			} else if inTypeOrAct == "bird" {
				animalSlice = append(animalSlice, Bird{name: inNameOrAnimal})
				fmt.Println("New type of bird created")

			} else if inTypeOrAct == "snake" {
				animalSlice = append(animalSlice, Snake{name: inNameOrAnimal})
				fmt.Println("New type of reptile created")

			} else {
				fmt.Println("That is not a valid animal option")
			}

		case "query":
			for _, animal := range animalSlice {

				if animal.getName() == inNameOrAnimal {
					if inTypeOrAct == "eat" {
						animal.Eat()

					} else if inTypeOrAct == "move" {
						animal.Move()

					} else if inTypeOrAct == "speak" {
						animal.Speak()

					} else {
						fmt.Println("That is not a valid activity")
					}
				}
			}
		}
	}
}
