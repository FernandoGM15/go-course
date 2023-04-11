package main

import (
	"fmt"
	"strings"
)

// Closures: function that returns another function
func repeat(number int) func(word string) string {
	return func(word string) string {
		return strings.Repeat(word, number)
	}
}
func main() {
	// Anonymous function (Like JS :D)
	func() {
		fmt.Println("Hello from anonymous functions")
	}()

	// Anonymous function assigned into variable
	myFunction := func(name string) string {
		return fmt.Sprintf("Hello %s, i'm a anonymous function in a variable", name)
	}
	fmt.Println(myFunction("Fer"))

	//Closures
	repeat5 := repeat(5)
	fmt.Println(repeat5("Hello"))

	repeat3 := repeat(3)
	fmt.Println(repeat3("Fer"))
}
