package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	var limit = rate.Every(200 * time.Millisecond)
	var limiter = rate.NewLimiter(limit, 3)
	for i := 0; i < 15; i++ {
		log.Printf("got #%d,err:%v", i, limiter.Wait(context.Background()))
	}

}
