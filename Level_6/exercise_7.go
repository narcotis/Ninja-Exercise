package main

import "fmt"

func main() {

	f := func(x int) {
		total := 0
		for i := 0; i < x; i++ {
			total += i
		}
		fmt.Println(total)
	}

	f(7)

}
