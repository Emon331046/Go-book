package main

import "fmt"

var x string = "hello world global"

func main() {
	var (
		x     = "hello world"
		z int = 5
	)
	fmt.Println(x)
	fmt.Println(x)
	x = x + " third"
	fmt.Println(x)
	y := "ola"

	fmt.Println(y)
	z = 5
	var a int
	var err error
	_, err = fmt.Scanf("%d%d", &z, &a)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(z, a)
	fun()

}
func fun() {
	fmt.Println(x)

}
