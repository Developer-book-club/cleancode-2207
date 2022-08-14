package shape_oop

import "cleancode-ex/point"

type Shape interface {
	CenterPoint() point.Point
	Area() int
}

type Square struct {
	TopLeft point.Point
	Side    int
}

func (s *Square) Area() int {
	return s.Side * s.Side
}

func (s *Square) CenterPoint() point.Point {
	p := point.New()
	p.SetCartesian(s.TopLeft.X()+s.Side/2, s.TopLeft.Y()+s.Side/2)

	return p
}

type Rectangle struct {
	TopLeft point.Point
	Height  int
	Width   int
}

func (r *Rectangle) Area() int {
	return r.Height * r.Width
}

func (r *Rectangle) CenterPoint() point.Point {
	p := point.New()
	p.SetCartesian(r.TopLeft.X()+r.Width/2, r.TopLeft.Y()+r.Height/2)

	return p
}

type Circle struct {
	Center point.Point
	Radius int
}

func (c *Circle) Area() int {
	return int(float64(c.Radius) * float64(c.Radius) * 3.14)
}

func (c *Circle) CenterPoint() point.Point {
	return c.Center
}
