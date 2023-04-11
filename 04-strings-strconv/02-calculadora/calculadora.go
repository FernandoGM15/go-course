package main

import (
	"fmt"
	"strconv"
	"strings"
)

func sumar(expresion string) int {
	sumArray := strings.Split(expresion, "+")
	var result int
	for _, value := range sumArray {
		//Error handling
		num, error := strconv.Atoi(value)

		//Ya que la función Atoi devuelve un error en caso de error, se verifica si ese parámetro es de tipo nil
		if error != nil {
			fmt.Println(error)
			fmt.Println("Favor de ingresar un numero entero")
			fmt.Println("Solo se realizan operaciones con cáracter +")
		} else {
			result += num
		}
	}
	return result
}

func main() {
	var expresion string
	var result int

	fmt.Print("=>")
	fmt.Scanln(&expresion)

	result = sumar(expresion)

	fmt.Printf("El resultado de la expresión es: %d \n", result)
}
