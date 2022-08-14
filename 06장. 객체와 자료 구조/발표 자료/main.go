package main

import (
	point1 "cleancode-ex/point"
	point2 "cleancode-ex/point_concrete"
	shape_oop "cleancode-ex/shape-oop"
	shape_pp "cleancode-ex/shape-pp"
	"fmt"
)

func pointExample() {
	p := point1.New()
	p.SetCartesian(2, 2)
	fmt.Println(p.X(), p.Y())
}

func pointConcreteExample() {
	p := &point2.Point{
		X: 2,
		Y: 2,
	}
	fmt.Println(p.X, p.Y)
}

func shapeOOPExample() {
	s := &shape_oop.Square{
		TopLeft: point1.New(),
		Side:    2,
	}
	fmt.Println(s.Area())
	fmt.Println(s.CenterPoint())
}

func shapePPExample() {
	s := &shape_pp.Square{
		TopLeft: point1.New(),
		Side:    2,
	}
	fmt.Println(shape_pp.Area(s))
	fmt.Println(shape_pp.CenterPoint(s))
}

func main() {
	pointExample()
	pointConcreteExample()
	shapeOOPExample()
	shapePPExample()
}
