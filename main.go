package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	var p sync.Pool

	p.New = func() any {
		return &http.Client{
			Timeout: 5 * time.Second,
		}
	}

	var wg sync.WaitGroup
	wg.Add(10)
	go func() {

		for i := 0; i < 10; i++ {
			go func() {
				defer wg.Done()

				c := p.Get().(*http.Client)
				defer p.Put(c)
				resp, err := c.Get("https://www.baidu.com")
				if err != nil {
					fmt.Println("failed to get baidu.com:", err)
					return
				}

				err = resp.Body.Close()
				if err != nil {
					fmt.Println("close response error:", err)
					return
				}
				fmt.Println("get baidu.com")

			}()
		}
	}()
	wg.Wait()

}
