package square

type Point struct {
	x, y int
}
func NewPoint(x, y int) Point{
	return Point{x, y}
}

type Square struct {
	start Point
	a     uint
}

func (s *Square) End() Point {
	return NewPoint(s.start.x + int(s.a), -s.start.y)
}

func (s *Square) Area() uint {
	return s.a * uint(s.start.y * 2)
}

func (s *Square) Perimeter() uint {
	return s.a * 2 + uint(s.start.y) * 2
}