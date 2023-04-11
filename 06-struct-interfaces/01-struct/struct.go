package main

import "fmt"

// Struct
type Persona struct {
	nombre string
	edad   int
}

// Definición de métodos de "clase"
func (p *Persona) print() {
	fmt.Printf("Nombre: %s \nEdad: %d \n===========\n", p.nombre, p.edad)
}

func (p *Persona) setName(name string) {
	p.nombre = name
}

func main() {
	// Inicialización de struct (clase)
	p1 := Persona{"Fernando", 24}
	p1.print()
	// Modificación de un atributo definido
	// p1.nombre = "Fer"
	p1.setName("Fer")
	p1.print()

	// Inicialización de struct, asignación key:value
	p2 := Persona{
		nombre: "Vane",
		edad:   24,
	}
	p2.print()

	p2.setName("Vanessa")
	p2.print()
}
