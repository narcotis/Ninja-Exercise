package main

import "fmt"

type hotdog int

var x4 hotdog
var y3 int

func main() {
	fmt.Printf("%v\n", x4)
	fmt.Printf("%T\n", x4)
	x4 = 42
	fmt.Printf("%v\n", x4)

	y3 = int(x4)
	fmt.Println(y3)
	fmt.Printf("%T\n", y3)
}
