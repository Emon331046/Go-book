package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Println("Hello "+"world")

	fmt.Println(len("Hello "+"world"))
	fmt.Println("hello 1+1 = \n" , 1.0+1.0)
	fmt.Println("hello World"[1])
	os.Exit(0)
}