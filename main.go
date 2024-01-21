package main

import (
	"log"
	"time"

	"go.uber.org/ratelimit"
)

func main() {
	rl := ratelimit.New(1, ratelimit.WithSlack(100)) // per second, no slack.

	for i := 0; i < 200; i++ {
		t := rl.Take()
		log.Println(t)
		log.Printf("got #%d", i)
		if i == 3 {
			time.Sleep(3 * time.Second)
		}
	}
}
