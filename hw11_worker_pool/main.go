package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := counter(500)
	fmt.Println(counter)
}

func counter(count int) (value int) {
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i := 0; i < count; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			value++
			mu.Unlock()
		}()
	}

	wg.Wait()

	return value
}
