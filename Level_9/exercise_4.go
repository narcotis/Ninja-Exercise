package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	counter := 0
	wg.Add(100)

	for i := 0; i < 100; i++ {
		go func() {
			mu.Lock()
			t := counter
			t++
			counter = t
			mu.Unlock()
			runtime.Gosched()
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(counter)
}
