package main

import "fmt"

func suma(a, b int) int {
	return a + b
}

func main() {
	var a int
	var b int
	fmt.Print("Primer valor: ")
	fmt.Scanln(&a)
	fmt.Print("Segundo valor: ")
	fmt.Scanln(&b)

	result := suma(a, b)
	fmt.Println("El resultado de la suma es: ", result)
}
