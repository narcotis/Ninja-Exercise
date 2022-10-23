package main

import "fmt"

type person struct {
	first    string
	last     string
	favorite []string
}

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

	fmt.Println(p1)
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.favorite {
		fmt.Println(i, v)
	}

	fmt.Println(p2)
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.favorite {
		fmt.Println(i, v)
	}

}
