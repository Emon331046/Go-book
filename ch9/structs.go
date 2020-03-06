package main

import (
	"fmt"
	"math"
)

type circle struct {
	x float64
	y float64
	r float64
}

func circleArea(c *circle) float64 {
	return math.Pi * c.r * c.r
}
func main() {

	//var c circle
	//c:= new(circle)
	//c := circle{x: 0, y: 0, r: 5}
	c := circle{0, 0, 5}
	fmt.Println(circleArea(&c))
}
