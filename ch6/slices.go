package main

import "fmt"

func main() {
	slice1 := make([]float64, 5)
	slice1[1] = 6
	fmt.Println(slice1)
	slice2 := make([]float64, 5, 10)
	fmt.Println(slice2)
	x := [10]int{
		10,
		11,
		12,
		13,
		14,
		15,
		16,
		17,
		18,
		19,
	}
	slicex := x[5:9]
	println(slicex) //output :[4/5]0xc00007a078
	// total slice size/ starting pointer number and address

	slicex0_4 := x[:5]
	fmt.Println(slicex0_4)

	slicex5_9 := x[5:]
	fmt.Println(slicex5_9)

	slicex0_9 := x[:]
	fmt.Println(slicex0_9)
	//append
	sliceappend1_2 := append(slicex0_4, 1, 2, 3)
	fmt.Println(sliceappend1_2, " another ", slicex)
	//copy

	sl1 := []int{1, 2, 3, 4}
	sl2 := make([]int, 5)
	copy(sl2, sl1)
	fmt.Println(sl2)
	sl3 := []int{1, 2, 3, 4}
	sl4 := make([]int, 3)
	copy(sl4, sl3)
	fmt.Println(sl4)
	//copy sl3 in sl4

}
