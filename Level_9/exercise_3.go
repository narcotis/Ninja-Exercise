package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	counter := 0
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			t := counter
			runtime.Gosched()
			t++
			counter = t
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
