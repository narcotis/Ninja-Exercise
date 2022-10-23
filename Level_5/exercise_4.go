package main

import "fmt"

func main() {

	an := struct {
		first     string
		friends   map[string]int
		favorites []string
	}{
		first: "Narco",
		friends: map[string]int{
			"Miller": 30,
			"Ben":    28,
		},
		favorites: []string{
			"Meat", "Vanilla", "Coffee",
		},
	}

	fmt.Println(an)
	for k, v := range an.friends {
		fmt.Println(k, v)
	}
}
