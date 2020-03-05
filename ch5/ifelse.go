package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {

		if i%2 == 0 && i%3 == 0 {
			fmt.Println(i, "divide by 2 & 3")
		} else if i%2 == 0 {
			fmt.Println(i, "divide by 2")
		} else if i%3 == 0 {
			fmt.Println(i, "divide by 3")
		} else {
			fmt.Println(i, "not divide by 2 & 3")
		}
	}

}
