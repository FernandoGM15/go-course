package figuras

import "math"

type Circle struct {
	Radius float32
}

func (c *Circle) Area() float32 {
	a := math.Pi * (c.Radius * c.Radius)
	return a
}

func (c *Circle) Perimeter() float32 {
	p := 2 * math.Pi * c.Radius
	return p
}
