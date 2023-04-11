package main

import "fmt"

func main() {
	// Arrays: Se coloca el tamaño que tendrá entre paréntesis seguido del tipo de datos/
	var numeros [5]int
	fmt.Println(numeros)

	numeros[0] = 10
	numeros[1] = 20
	numeros[2] = 30

	fmt.Println(numeros)
	fmt.Println(numeros[4])

	//Arrays con valores
	nombres := [3]string{"hola", "como", "estas"}
	fmt.Println("Nombres: ", nombres)

	//Arrays con valores sin longitud
	colores := [...]string{"amarillo", "azul", "rojo", "morado", "verde"}
	fmt.Println("Colores: ", colores, "Con tamaño: ", len(colores))

	//Arrays con indices
	monedas := [...]string{
		0: "Dolar",
		2: "Soles",
		3: "Pesos",
		5: "Euro",
	}
	fmt.Println("Monedas: ", monedas, "Con tamaño: ", len(monedas))

	//Sub array NOTA: Se obtendrán solo 2 elementos ya que la lógica plantea obtener los elementos antes del 3 elemento
	// subMoneda := monedas[0:3]
	subMoneda := monedas[3:]
	fmt.Println("Submoneda: ", subMoneda, "Con tamaño: ", len(subMoneda))
}
