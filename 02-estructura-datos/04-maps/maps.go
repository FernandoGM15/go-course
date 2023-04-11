package main

import (
	"fmt"
)

func main() {
	//Inicialización de map
	dias := make(map[int]string)
	fmt.Println(dias)

	//Agregar datos
	dias[1] = "Domingo"
	dias[2] = "Lunes"
	dias[3] = "Martes"
	dias[4] = "Miércoles"
	dias[5] = "Jueves"
	dias[6] = "Viernes"
	dias[7] = "Sábado"
	fmt.Println(dias)

	//Sobrescribir datos
	dias[7] = "SÁBADO"
	fmt.Println(dias)

	//Eliminar datos
	delete(dias, 1)
	fmt.Println(dias)

	//Nuevo Map
	estudiantes := make(map[string][]int)

	estudiantes["Alex"] = []int{13, 15, 16}
	estudiantes["Roel"] = []int{14, 13, 17}

	fmt.Println(estudiantes)
	fmt.Println(estudiantes["Alex"][1])
}
