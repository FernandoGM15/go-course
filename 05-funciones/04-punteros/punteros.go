package main

import "fmt"

func modifyValue(word *string) {
	fmt.Printf("%p\n", word)
	*word = "Hello from function"
}

func main() {
	word := "Hello from go"
	fmt.Printf("%p\n", &word)
	fmt.Println("Before the function:", word)

	modifyValue(&word)

	fmt.Println("After the function:", word)
}
