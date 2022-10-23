package main

import "fmt"

var x int = 324

func main() {
	fmt.Printf("%d\t%b\t%#x\n", x, x, x)

	x = x << 1
	fmt.Printf("%d\t%b\t%#x", x, x, x)
}
