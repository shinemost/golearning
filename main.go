package main

import (
	"fmt"
	"net/http"
	"runtime"
	"sync"
	"time"
)

func main() {
	var p sync.Pool

	for i := 0; i < 10; i++ {
		p.Put(&http.Client{
			Timeout: 5 * time.Second,
		})
	}
	runtime.GC()
	runtime.GC()

	c := p.Get().(*http.Client)
	req, err := c.Get("https://www.baidu.com")
	if err != nil {
		fmt.Println("failed to get baidu")
	}

	err = req.Body.Close()
	if err != nil {
		return
	}
}
