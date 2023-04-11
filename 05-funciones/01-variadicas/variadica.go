package main

import "fmt"

/*
* The functions receive and return multiples parameters
* On receive:  It needs to have "..." before the type of parameter, They must be at the end of the params
* On return: It needs to have the return types must be inside of parentheses
* the params is treated as an array
 */
func sum(name string, numbers ...int) (string, int) {
	mensaje := fmt.Sprintf("Hola %s", name)
	result := 0
	for _, item := range numbers {
		result += item
	}
	return mensaje, result
}

func main() {
	message, result := sum("Fernando", 10, 20, 30)
	fmt.Printf("%s, El resultado es %d \n", message, result)
}
