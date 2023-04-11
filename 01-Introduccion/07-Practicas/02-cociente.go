package main

import "fmt"

func cociente(a, b int) int {
	return a / b
}

func residuo(a, b int) int {
	return a % b
}

func main() {
	var a, b int
	fmt.Print("Ingrese el dividendo: ")
	fmt.Scanln(&a)

	fmt.Print("Ingrese el divisor: ")
	fmt.Scanln(&b)

	cociente := cociente(a, b)
	residuo := residuo(a, b)

	fmt.Println("El cociente de la division es: ", cociente)
	fmt.Println("El residuo de la division es: ", residuo)
}
