package main

import "fmt"

func one(xptr1 *int) {
	*xptr1 = 1
}
func swap(xptr *int, yptr *int) {
	var temp *int
	temp = xptr
	xptr = yptr
	yptr = temp
	fmt.Println(*xptr, *yptr)
	tm := *xptr
	*xptr = *yptr
	*yptr = tm
}
func main() {
	//xptr := new(int)
	//one(xptr)
	//fmt.Println(xptr)
	//fmt.Println(*(&xptr))
	//fmt.Println(*xptr)
	x := 5
	y := 6
	swap(&x, &y)
	fmt.Println(x, y)

}
