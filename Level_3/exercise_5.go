package main

import "fmt"

func main() {
	for num := 10; num <= 100; num++ {
		fmt.Println(num, num/4, num%4)
	}
}
