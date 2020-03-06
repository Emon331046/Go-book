package main

import "fmt"

type Person struct {
	Name string
}
type Android struct { // android subclass
	Person //person super class
	Model  string
}

func (p *Person) Talk() { //ei method pabe na override hbe
	fmt.Println("Hi, my name is person", p.Name)
}
func (p *Android) Talk() {
	fmt.Println("Hi, my name androiid", p.Name)
}
func main() {
	a := new(Android)
	a.Talk()
}
