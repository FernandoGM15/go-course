package main

import "fmt"

func main() {
	a := 20
	b := 10

	//Sumatoria
	result := a + b
	fmt.Println("Suma: ", result)

	//Resta
	result = a - b
	fmt.Println("Resta: ", result)

	//Multiplicación
	result = a * b
	fmt.Println("Multiplicación: ", result)

	//Division
	result = a / b
	fmt.Println("Division: ", result)

	//Division
	var div float32 = 3.0 / 2.0
	fmt.Println("Division exacta: ", div)

	//Modulo
	result = 3 % 2
	fmt.Println("Modulo: ", result)

}
