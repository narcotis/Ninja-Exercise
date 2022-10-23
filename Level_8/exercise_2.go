package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string `json:"first,omitempty"`
	Last  string `json:"last,omitempty"`
	Age   int    `json:"age,int"`
}

type People []Person

func main() {
	ps := `[{"first":"Narco","last":"Yun","age":32},{"first":"James","last":"Bond","age":33}]`

	px := People{}

	err := json.Unmarshal([]byte(ps), &px)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(px)
}
