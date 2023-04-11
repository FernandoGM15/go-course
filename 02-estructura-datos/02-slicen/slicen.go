package main

import "fmt"

func main() {
	//Slicen:
	numeros := []int{1, 2, 3}
	fmt.Println(numeros, len(numeros))

	//Agregar elementos
	numeros = append(numeros, 4)
	numeros = append(numeros, 5)
	fmt.Println(numeros, len(numeros))

	//Sub Slicen: Se obtiene en referenciando a un array padre,
	// esa es la razón por la que si se modifica un valor en el arreglo padre,
	// también se modifica en el slicen
	subSlicen := numeros[:2]
	numeros[0] = 100

	fmt.Println(subSlicen)
	fmt.Println(numeros)

	meses := []string{"enero", "febrero", "marzo"}
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)

	meses = append(meses, "Abril")
	fmt.Printf("Len: %v, Cap: %v, %p \n", len(meses), cap(meses), meses)
}
