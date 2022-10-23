package main

import "fmt"

func main() {
	a := incrementor()
	b := incrementor()

	fmt.Println(a(), a(), a(), a())
	fmt.Println(b(), b(), b())
}

func incrementor() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
