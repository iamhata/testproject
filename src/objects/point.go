package objects

type IPoint interface {
	SetX( int)
	SetY( int)
}

type Point struct {
	X, Y int
}

var instance *Point

func GetInstance() *Point{
	if instance == nil {
		instance = new(Point)
	}
	return instance
}

func (s *Point) SetX(x int) {
	s.X = x;
}

func (s *Point) SetY(y int) {
	s.Y = y;
}