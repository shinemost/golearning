package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {

	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(3)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			err := sem.Acquire(context.Background(), 1)
			if err != nil {
				return
			}
			defer sem.Release(1)

			// Simulate some work
			fmt.Printf("Goroutine %d: Start\n", id)
			// Do some work...
			time.Sleep(3 * time.Second)
			fmt.Printf("Goroutine %d: End\n", id)
		}(i)
	}

	wg.Wait()
}
