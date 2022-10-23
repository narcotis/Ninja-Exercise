package main

import "fmt"

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))

}
