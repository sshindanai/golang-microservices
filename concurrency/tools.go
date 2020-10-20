package main

import (
	"fmt"
)

func main() {
	c := make(chan int, 3)

	fmt.Println("print 1")
	c <- 1

	fmt.Println("print 2")
	c <- 2

	fmt.Println("print 3")
	c <- 3

	fmt.Println("print 4")
	c <- 4
}

func helloWorld() {
	fmt.Println("Hello World")
}
