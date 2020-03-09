package main

import (
	"fmt"
	"os"
	"strconv"
)

func pinger(c chan string) {
	for i := 0; ; i++ {
		c <- "ping" + strconv.Itoa(i)
		if i == 5000 {
			c <- "hello"
		}
	}
}

//func ponger(c chan string) {
//	for i := 0; ; i++ {
//		c <- "pong"
//
//	}
//}
//A channel that doesn't have these restrictions
//is known as bi-directional. A bi-directional
//channel can be passed to a function that takes
//send-only or receive-only channels, but the
//reverse is not true.
func printer(c chan string) {
	for {
		msg := <-c
		fmt.Println(msg)
		//time.Sleep(time.Second * 1)
	}
}
func main() {
	var c chan string = make(chan string)
	go pinger(c)
	go printer(c)
	//var input string
	//fmt.Scanln(&input)

	for {
		if <-c == "hello" {
			fmt.Println("break it")
			os.Exit(0)
		}
	}
}
