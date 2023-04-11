package main

import "fmt"

func main() {
	hola := "Hola"
	mundo := "Mundo"

	//Se imprime todo en una sola linea
	fmt.Print(hola)
	fmt.Print(mundo)

	// Se imprime con salto de linea
	fmt.Println(hola)
	fmt.Println(mundo)

	nombre := "Fernando"
	edad := 28

	/* Se formatea el mensaje
	%s: string
	%d: number
	%v: any
	%T: Obtener tipo de variable
	*/
	fmt.Printf("Hola, me llamo %s y mi edad es %d años\n", nombre, edad)

	// Se almacena un mensaje formateado en una variable
	mensaje := fmt.Sprintf("Hola, me llamo %s y mi edad es %d años", nombre, edad)
	fmt.Println(mensaje)

	fmt.Printf("Nombre: %T \n", edad)

	//Almacenar input
	// &[variable] Se almacena el input en variable
	fmt.Print("Ingrese otro nombre: ")
	fmt.Scanln(&nombre)
	fmt.Println(nombre)

}
