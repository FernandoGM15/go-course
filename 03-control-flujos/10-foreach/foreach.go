package main

import "fmt"

func main() {
	nombres := [...]string{"Alex", "Roel", "Juan"}

	// for i := 0; i < len(nombres); i++ {
	// fmt.Println(nombres[i])
	// }
	for key, elemento := range nombres {
		fmt.Println(key, elemento)
	}
}
