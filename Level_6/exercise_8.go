package main

import "fmt"

func main() {

	e := even(1, 2, 3, 4, 5, 6, 7, 8, 9)
	e()
}

func even(xi ...int) func() {
	var ii []int
	for _, v := range xi {
		if v%2 == 0 {
			ii = append(ii, v)
		}
	}
	return func() {
		fmt.Println(ii)
	}
}
