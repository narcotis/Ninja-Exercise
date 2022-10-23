package main

import "fmt"

func main() {
	defer close()
	fmt.Println("Main started")

}

func close() {
	defer func() {
		fmt.Println("close DEFER")
	}()
	fmt.Println("Main closed")
}
