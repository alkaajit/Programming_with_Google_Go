package main

import "fmt"

type animal struct {
	food,locomotion,sound string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (an animal) Eat() {
	fmt.Println(an.food)
	return
}

func (an animal) Move() {
	fmt.Println(an.locomotion)
	return
}

func (an animal) Speak() {
	fmt.Println(an.sound)
	return
}

func main() {
	animalMap := make(map[string]animal)
	animalMap["cow"] = animal{"grass", "walk", "moo"}
	animalMap["bird"] = animal{"worms", "fly", "peep"}
	animalMap["snake"] = animal{"mice", "slither", "hsss"}
	
	var generalAnimal animalInterface
	
	for {
		var command, animalType, animalData string
		fmt.Print(">")
		fmt.Scan(&command, &animalType, &animalData)
		if command == "query" {
			generalAnimal = animalMap[animalType]
			switch animalData {
			case "eat":
				generalAnimal.Eat()
			case "move":
				generalAnimal.Move()
			case "speak":
				generalAnimal.Speak()
			}
		}
		if command == "newanimal" {
			animalMap[animalType] = animalMap[animalData]
			fmt.Println("Created it!")
		}
	}
}
