package main

import "fmt"

var s = `Raw string variable`

func main() {
	fmt.Printf("%s\t%T\n", s, s)
}
