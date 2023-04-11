package main

import "fmt"

// Estructura Padre
type Persona struct {
	name string
	age  int
}

func (p *Persona) print() {
	fmt.Println("-------------------")
	fmt.Printf("Name: %s\nAge: %d\n", p.name, p.age)
	fmt.Println("-------------------")

}

// Herencia de la estructura padre
type Empleado struct {
	Persona
	sueldo float32
}

func main() {

	empleado1 := Empleado{
		sueldo: 500.52,
	}
	empleado1.name = "Pedro"
	empleado1.age = 30

	// empleado1.print()
	fmt.Println(empleado1)
}
