package main

import "fmt"

func factory(n int) int {
	if n == 0 {
		return 1
	}
	return n * factory(n-1)
}

func main() {
	fmt.Println(factory(5))
}
