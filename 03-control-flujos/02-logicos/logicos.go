package main

import "fmt"

func main() {
	//Not: Devuelve el valor opuesto
	// fmt.Println(!true)

	//And: Ambos valores deben ser true
	// fmt.Println(false && false)

	//Or: Al menos un valor debe ser true
	// fmt.Println(false || false)

	b := 2
	r := b == 2 && !(4 > b)
	fmt.Println(r)
}
