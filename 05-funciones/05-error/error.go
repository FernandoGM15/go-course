package main

import (
	"errors"
	"fmt"
)

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No es posible division entre 0")
	}
	return dividendo / divisor, nil
}

func main() {
	res, err := division(4, 0)
	if err == nil {
		fmt.Println("Resultado: ", res)
	} else {
		fmt.Println(err)
	}
}
