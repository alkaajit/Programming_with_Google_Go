package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, sound string
}

func (v Animal) Eat() {
	fmt.Println(v.food)
}

func (v Animal) Move() {
	fmt.Println(v.locomotion)
}

func (v Animal) Speak() {
	fmt.Println(v.sound)
}

func main() {
	animal_map := map[string]Animal{
		"cow" : Animal{"grass","walk","moo"},
		"bird" : Animal{"worms","fly","peep"},
		"snake" : Animal{"mice","slither","hsss"},
	}
	for{
		fmt.Print(">")
		name :="0"
		action :="0"
		fmt.Scan(&name,&action)
		if action=="eat"{
			animal_map[name].Eat()
		} else if action=="move"{
			animal_map[name].Move()
		} else if action=="speak"{
			animal_map[name].Speak()
		} else{
			fmt.Println("Invalid input. Try again")
		}
	}
}
