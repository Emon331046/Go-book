package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		switch i % 2 {
		case 0:
			fmt.Println(i, "even")
		default:
			fmt.Println(i, "odd")

		}
	}

}
