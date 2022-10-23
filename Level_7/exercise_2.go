package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}

func changeMe(p *person) {
	(*p).age += 1
}

func main() {
	p := person{
		first: "Narco",
		last:  "Yun",
		age:   32,
	}

	fmt.Println(p)
	changeMe(&p)
	fmt.Println(p)

}
