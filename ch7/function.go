package main

import "fmt"

func main() {
	//var x int = 5
	//f(x)
	//println(&x)
	//a, b, c, d := f2()
	//fmt.Println(a, b, c, d)
	//fmt.Println(f3(1, 2, 3, 4, 5))
	//xs := []int{1, 2, 3, 4, 5}
	//fmt.Println(f4(xs...))
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	//
	//increment := func(x int) int {
	//	x++
	//	return x
	//}
	//fmt.Println(increment(5))

}

//take input
func f(x int) {
	//println(&x)
	fmt.Println(x)
}

//returning multiple variable
func f2() (int, int, string, []int) {
	x := []int{
		6, 7, 8,
	}
	return 5, 7, "emon", x
}
func f3(arg ...int) int {
	total := 0
	for _, v := range arg {
		total += v
	}
	return total
}
func f4(arg ...int) int {
	total := 1
	for _, v := range arg {
		total *= v
	}
	return total
}

func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}
