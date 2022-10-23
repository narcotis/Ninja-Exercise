package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("Not printed")
	case true:
		fmt.Println("Printed")
	default:
		fmt.Println("Default")
	}
}
