package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/marusama/semaphore/v2"
)

func main() {
	sem := semaphore.New(3)
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {

		go func(id int) {
			defer wg.Done()

			err := sem.Acquire(context.Background(), 1)
			if err != nil {
				return
			} // Acquire semaphore

			//count := sem.GetCount()
			//println(count)
			sem.SetLimit(4)
			println(sem.GetCount())
			defer sem.Release(1) // Release semaphore

			// Simulate some work
			fmt.Printf("Goroutine %d: Start\n", id)
			// Do some work...
			time.Sleep(3 * time.Second)
			fmt.Printf("Goroutine %d: End\n", id)
		}(i)
	}

	wg.Wait()
}
