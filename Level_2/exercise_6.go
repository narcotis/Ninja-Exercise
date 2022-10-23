package main

import "fmt"

const (
	_ = iota
	i = 2017 + iota
	j = 2017 + iota
	k = 2017 + iota
	l = 2017 + iota
)

func main() {
	fmt.Println(i, j, k, l)
}
