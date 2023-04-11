package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		//Salto a la siguiente iteración
		if i == 5 {
			fmt.Println("Salta la iteración")
			continue
		}
		//Rompe el bucle
		if i == 8 {
			fmt.Println("Romper bucle")
			break
		}
		fmt.Println(i)
	}
}
