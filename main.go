package main

import (
	"fmt"

	"github.com/sourcegraph/conc"
)

func main() {
	var wg conc.WaitGroup
	for i := 0; i < 10; i++ {
		i := i
		wg.Go(func() {
			//m := 0
			//n := i / m
			//fmt.Println(n)
			fmt.Println(i)
		})
	}
	//wg.Wait()
	wg.WaitAndRecover()

}

/*func main() {
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		i := i
		go func() {
			defer wg.Done()
			m := 0
			n := i / m
			fmt.Println(n)
		}()

	}
	wg.Wait()
}*/
