package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Age   int    `json:"age,int"`
}

func main() {
	p1 := person{
		First: "Narco",
		Last:  "Yun",
		Age:   32,
	}

	p2 := person{
		First: "James",
		Last:  "Bond",
		Age:   33,
	}

	px := []person{p1, p2}
	fmt.Println(px)

	bs, err := json.Marshal(px)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
