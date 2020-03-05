package main

import "fmt"

func main() {
	//	fmt.Println(`1
	//2
	//3
	//4
	//5
	//6
	//7
	//8
	//9
	//10`)

	//i := 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i++
	//}
	//for i := 1; i <= 10; i++ {
	//
	//	if i%2 == 0 {
	//		fmt.Println(i, "even")
	//	} else {
	//		fmt.Println(i, "odd")
	//	}
	//}
	for i := 1; i <= 10; i++ {
		switch i % 2 {
		case 0:
			fmt.Println(i, "even")
		default:
			fmt.Println(i, "odd")

		}
	}

}
