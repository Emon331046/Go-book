package main

import (
	"fmt"
	"math"
)

type circle1 struct {
	x float64
	y float64
	r float64
}

func (c *circle1) circleArea1() float64 {
	return math.Pi * c.r * c.r
}
func main() {

	//var c circle
	//c:= new(circle)
	//c := circle{x: 0, y: 0, r: 5}
	c := circle1{0, 0, 5}
	fmt.Println(c.circleArea1())
}
