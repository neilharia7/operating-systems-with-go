package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
	runtime.GOMAXPROCS(4)
	value, x := 0, 5
	var wg sync.WaitGroup
	var mu sync.Mutex

	// annonymous function
	decrement := func() {
		defer wg.Done()
		mu.Lock()
		value -= x
		mu.Unlock()
	}

	// annonymous function
	increment := func() {
		defer wg.Done()
		mu.Lock()
		value += x
		mu.Unlock()
	}

	for i :=0; i < 200; i++ {
		wg.Add(2)

		go increment()
		go decrement()
	}

	wg.Wait()
	fmt.Print(value)
}