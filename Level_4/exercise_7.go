package main

import "fmt"

func main() {

	a := []string{"James", "Bond", "Shaken, not stirred"}
	b := []string{"Miss", "Moneypenny", "Helloooooo, James."}

	s := [][]string{a, b}
	for i, v := range s {
		fmt.Println(i, v)
	}

	for i, xxs := range s {
		for j, xs := range xxs {
			fmt.Println(i, j, xs)
		}
	}
}
