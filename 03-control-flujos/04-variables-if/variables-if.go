package main

import "fmt"

func main() {
	if nombre, edad := "Roel", 24; nombre == "Alex" {
		fmt.Println("Hola, ", nombre)
	} else {
		fmt.Printf("Nombre: %s - Edad: %d \n", nombre, edad)
	}

	users := make(map[string]string)
	users["Alex"] = "alex@gmail.com"
	users["Roel"] = "roel@gmail.com"

	// correo, ok := users["Juan"]
	if correo, ok := users["Roel"]; ok {
		fmt.Println(correo)
	} else {
		fmt.Println("No fue posible obtener el valor")
	}
}
