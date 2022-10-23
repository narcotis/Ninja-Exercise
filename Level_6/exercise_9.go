package main

import "fmt"

func odd(f func(xi ...int) int, xxi ...int) int {
	var ii []int
	for _, v := range xxi {
		if v%2 != 0 {
			ii = append(ii, v)
		}
	}
	return f(ii...)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func main() {
	fmt.Println(odd(sum, 1, 2, 3, 4, 5, 6, 7, 8, 9))
}
