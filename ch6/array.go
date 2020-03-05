package main

import "fmt"

func main() {
	x := [5]int{98,
		93,
		77,
		//82,
		//83,
	}
	//x[4] = 100
	//x[0] = 98
	//x[1] = 98
	//x[2] = 101
	//x[3] = 102
	fmt.Println(x)
	var total float64
	//for i := 0; i < len(x); i++ {
	//	total += float64(x[i])
	//}
	//for i, value := range x {
	//	total += float64(value)
	//}
	for _, value := range x {
		total += float64(value)
	}
	fmt.Println(total)
	fmt.Println(total / float64(len(x)))

}
