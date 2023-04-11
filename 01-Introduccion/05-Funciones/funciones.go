package main

import "fmt"

func saludar(nombre string, edad int) {
	fmt.Printf("Hola me llamo %s y tengo %d años\n", nombre, edad)
}

func sumar(a, b int) int {
	return a + b
}

func main() {
	saludar("Fernando", 24)
	suma := sumar(10, 30)
	fmt.Println(suma)
}
