package main

import (
	"log"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	rl := ratelimit.New(1, ratelimit.WithSlack(3))

	for i := 0; i < 10; i++ {
		t := rl.Take()
		log.Printf("got #%d,get time:%s", i, t.Format(time.RFC3339))
		if i == 3 {
			time.Sleep(3 * time.Second)
		}
	}
}
