package main

import "fmt"

type narco int

var x3 narco

func main() {
	fmt.Printf("%v\n", x3)
	fmt.Printf("%T\n", x3)
	x3 = 42
	fmt.Printf("%v\n", x3)
}
