package main

import (
	"fmt"
)

func main() {
	g := 12 == 32
	h := 42 <= 21
	i := 42 >= 21
	j := 22 != 21
	k := 22 < 21
	l := 22 > 21

	fmt.Println(g, h, i, j, k, l)
	fmt.Printf("%d\t%d\t%d\t%s\t%s\t%s\n", g, h, i, j, k, l)
}
