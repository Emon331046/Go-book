package main

import (
	"fmt"
	"math"
)

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

type rectangle struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func (r *rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

type circle2 struct {
	x float64
	y float64
	r float64
}

func (c *circle2) area() float64 {
	return math.Pi * c.r * c.r
}

type Shape interface {
	area() float64
}

func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func main() {

	//var c circle
	//c:= new(circle)
	//c := circle{x: 0, y: 0, r: 5}
	c := circle2{0, 0, 5}
	r := rectangle{0, 0, 5, 10}
	fmt.Println(c.area())
	fmt.Println(r.area())
	fmt.Println(totalArea(&c, &r))
}
