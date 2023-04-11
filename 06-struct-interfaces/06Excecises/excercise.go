package main

import (
	// "Curso/figuras"
	// "Curso/interfaces"
	"Curso/models"
	"fmt"
)

func main() {
	// square := figuras.Square{
	// 	Side: 4,
	// }
	// interfaces.GetArea(&square)

	// circle := figuras.Circle{
	// 	Radius: 5,
	// }
	// interfaces.GetArea(&circle)

	p1 := models.Person{}
	p1.Constructor("Fernando", 24)
	fmt.Println(p1)
	p1.SetName("Vane")
	fmt.Println(p1)

}
