package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var wg sync.WaitGroup
	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
	fmt.Println(count)
}

// it will be below 1,000,000
// What you have here is a race condition. The operation to increment an integer is not atomic, meaning the code can get interrupted mid-operation.
// To solve this problem, you can use a sync.Mutex to ensure only one goroutine updates count at a time.
