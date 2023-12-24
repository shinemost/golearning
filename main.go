package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var m sync.Map

	wg.Add(5)

	for i := 0; i < 5; i++ {
		i := i
		go func() {
			m.Store(i, fmt.Sprintf("test #%d", i))
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("store done")

	wg.Add(5)

	for i := 0; i < 5; i++ {
		i := i
		go func() {
			t, _ := m.Load(i)
			fmt.Println("for loop: ", t)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println("load done")

	m.Range(func(k, v any) bool {
		fmt.Println("range ():", v)
		return true
	})

}
