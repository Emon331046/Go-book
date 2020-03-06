package main

import "fmt"

func first() {
	fmt.Println("1st")
}
func third() {
	fmt.Println("third")
}
func fourth() {
	fmt.Println("4th")
}
func second() {
	fmt.Println("2nd")
}
func main() {
	first()
	first()
	defer third()
	first()
	defer second()
	defer fourth()
}
