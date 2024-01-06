package main

import (
	"fmt"
	"sync"
)

func main() {
	const concurrentAccess = 3

	var wg sync.WaitGroup
	sem := make(chan struct{}, concurrentAccess)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			sem <- struct{}{} // Acquire semaphore
			defer func() {
				<-sem // Release semaphore
			}()

			// Simulate some work
			fmt.Printf("Goroutine %d: Start\n", id)
			// Do some work...
			fmt.Printf("Goroutine %d: End\n", id)
		}(i)
	}

	wg.Wait()
}
