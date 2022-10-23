package main

import "fmt"

func main() {
	m := map[string][]string{
		"bond_james":      []string{"Shaken, not stirred", "Martinis", "Women"},
		"moneypenny_miss": []string{"James Bond", "Literature", "Computer Science"},
		"no_dr":           []string{"Being evil", "Ice cream", "Sunsets"},
	}

	m["fleming_ian"] = []string{"steaks", "cigars", "espionage"}

	for k, v := range m {
		fmt.Println("Record : ", k)
		for i, s := range v {
			fmt.Println("\tIndex : ", i, "\tValue : ", s)
		}
	}

}
