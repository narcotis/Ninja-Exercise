package main

import "fmt"

func main() {
	p1 := person{
		first: "Narco",
		last:  "Yun",
		favorite: []string{
			"Chocolate", "Vanilla", "Strawberry",
		},
	}

	p2 := person{
		first: "Miller",
		last:  "Kim",
		favorite: []string{
			"Vanilla", "Chocolate",
		},
	}

	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for k, v := range m {
		fmt.Println(k, v)
	}
}
