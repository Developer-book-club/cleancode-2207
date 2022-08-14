package shape_pp

import (
	"cleancode-ex/point"
	"cleancode-ex/point_concrete"
)

type Square struct {
	TopLeft point.Point
	Side    int
}

type Rectangle struct {
	topLeft point_concrete.Point
	height  int
	width   int
}

type Circle struct {
	center point.Point
	radius int
}

func CenterPoint(shape any) point.Point {
	x, y := 0, 0
	switch shape.(type) {
	case *Square:
		s := shape.(*Square)
		x = s.TopLeft.X() + s.Side/2
		y = s.TopLeft.Y() + s.Side/2
	case *Rectangle:
		r := shape.(*Rectangle)
		x = r.topLeft.X + r.width/2
		y = r.topLeft.Y + r.height/2
	case *Circle:
		c := shape.(*Circle)
		x = c.center.X()
		y = c.center.Y()
	}

	p := point.New()
	p.SetCartesian(x, y)

	return p
}

func Area(shape any) int {
	switch shape.(type) {
	case *Square:
		s := shape.(*Square)
		return s.Side * s.Side
	case *Rectangle:
		r := shape.(*Rectangle)
		return r.height * r.width
	case *Circle:
		c := shape.(*Circle)
		return int(float64(c.radius) * float64(c.radius) * 3.14)
	}

	return 0
}
