package main

import "fmt"

var messages string

func function1() {
	messages = "Hola desde función 1"
	fmt.Println(messages)
}

func function2() {
	messages = "Hola desde función 2"
	fmt.Println(messages)
}

func function3() {
	messages = "Hola desde función 3"
	fmt.Println(messages)
}

func main() {
	messages = "Hola desde la función principal"
	fmt.Println(messages)

	defer function1()
	function2()

	fmt.Println(messages)

}
