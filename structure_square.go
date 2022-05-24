package structure_square

type Point struct {
	x, y int
}

type Square struct {
	startPoint1, endPoint1, startPoint2, endPoint2 Point
}
func NewPoint(x, y int) Point{
	return Point{x, y}
}
func NewSquare(startPoint1 Point) (s Square) {
	s.startPoint1 = startPoint1
	s.endPoint1.x = s.startPoint1.x
	s.endPoint1.y = -s.startPoint1.y
	s.startPoint2.x = 2 * s.startPoint1.y
	s.startPoint2.y = s.startPoint1.y
	s.endPoint2.x = s.startPoint2.x
	s.endPoint2.y = -s.startPoint2.y
	return s
}

func (s *Square) Area() int {
	side := s.startPoint2.x - s.startPoint1.x
	return side * side
}

func (s *Square) Perimeter() int {
	side := s.startPoint2.x - s.startPoint1.x
	return side * 4
}
