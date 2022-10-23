package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name string
	Age  int
	sex  string
}

func main() {

	u1 := User{
		Name: "Narco",
		Age:  32,
		sex:  "Male",
	}

	u2 := User{
		Name: "Moneypenny",
		Age:  27,
		sex:  "Female",
	}

	ux := []User{u1, u2}
	err := json.NewEncoder(os.Stdout).Encode(ux)
	if err != nil {
		fmt.Println(err)
	}
}
