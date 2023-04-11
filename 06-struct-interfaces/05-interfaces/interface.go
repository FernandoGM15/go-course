package main

import "fmt"

// Interface for abstract classes
type Animal interface {
	move() string
}

type Dog struct{}

type Fish struct{}

type Bird struct{}

func (d *Dog) move() string {
	return "I'm a dog, i walk"
}

func (f *Fish) move() string {
	return "I'm a fish, i swim"
}

func (b *Bird) move() string {
	return "I'm a Bird, i fly"
}

func moveAnimal(animal Animal) {
	fmt.Println(animal.move())
}

func main() {
	dog := Dog{}
	moveAnimal(&dog)

	fish := Fish{}
	moveAnimal(&fish)

	bird := Bird{}
	moveAnimal(&bird)
}
