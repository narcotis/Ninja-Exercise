package main

import "fmt"

func main() {
	//var a [5]int
	//a[0] = 0
	//a[1] = 2
	//a[2] = 4
	//a[3] = 8
	//a[4] = 16

	a := [5]int{0, 2, 4, 8, 16}

	for i, v := range a {
		fmt.Println(i, v)
	}

	fmt.Printf("%T\n", a)
}
