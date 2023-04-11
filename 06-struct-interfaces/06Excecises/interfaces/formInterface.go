package interfaces

import "fmt"

// Interfaces only works with methods
type FormI interface {
	Area() float32
	Perimeter() float32
}

func GetArea(form FormI) {
	fmt.Println(form.Area())
	fmt.Println(form.Perimeter())
}
