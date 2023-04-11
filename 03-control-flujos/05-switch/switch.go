package main

import "fmt"

func main() {
	var char string
	fmt.Print("Ingrese una letra: ")
	fmt.Scanln(&char)

	// switch {
	// case char == "a" || char == "A":
	// 	fmt.Println(char, "es vocal")
	// case char == "e" || char == "E":
	// 	fmt.Println(char, "es vocal")
	// case char == "i" || char == "I":
	// 	fmt.Println(char, "es vocal")
	// case char == "o" || char == "O":
	// 	fmt.Println(char, "es vocal")
	// case char == "u" || char == "U":
	// 	fmt.Println(char, "es vocal")
	// default:
	// 	fmt.Println(char, "no es vocal")
	// }

	switch char {
	case "a", "A", "e", "E", "i", "I", "o", "O", "u", "U":
		fmt.Println(char, "es vocal")
	default:
		fmt.Println(char, "no es vocal")
	}

}
