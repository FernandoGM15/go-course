package figuras

type Square struct {
	Side float32
}

func (s *Square) Area() float32 {
	a := s.Side * s.Side
	return a
}

func (s *Square) Perimeter() float32 {
	p := 4 * s.Side
	return p
}
