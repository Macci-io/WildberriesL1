package main

import (
	"fmt"
	"math"
)

/*
Разработать программу нахождения расстояния между двумя точками,
которые представлены в виде структуры Point с инкапсулированными параметрами
x,y и конструктором.
*/
type Point struct {
	x, y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func (p *Point) Calc(an *Point) float64 {
	x := p.x - an.x
	y := p.y - an.y
	n := y*y + x*x
	return math.Sqrt(n)
}

func main() {
	a := NewPoint(10, 10)
	b := NewPoint(15, 15)
	fmt.Println(a.Calc(b))
}
