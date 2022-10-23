package main

import "fmt"

func main() {
	s := "Narco"
	if s == "James" {
		fmt.Println("This is not printed")
	} else if s == "Bond" {
		fmt.Println("This is not printed neither")
	} else {
		fmt.Printf("Hello %s\n", s)
	}
}
