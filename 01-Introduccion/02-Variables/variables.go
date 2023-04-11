package main

import "fmt"

func main() {
	// ====================Definición de variables
	var nombre1 string
	nombre1 = "Fernando"

	var nombre2 string = "Alex"

	//Shorthand
	edad := 24
	fmt.Println(nombre1, nombre2, edad)

	// ====================Valores default
	var a int     //Default: 0
	var b float64 //Default: 0
	var c string  //Default: ''
	var d bool    //Default: false
	fmt.Println(a, b, c, d)

	// ====================Definición de contantes
	const pi = 3.141592
	fmt.Println(pi)
}
