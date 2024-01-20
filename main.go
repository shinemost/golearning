package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)
	var limiter = rate.NewLimiter(1, 3)
	for i := 0; i < 3; i++ {
		log.Printf("got #%d,err:%v", i, limiter.Wait(context.Background()))
	}

	log.Println("set new limit at 10s")

	//设置新速率为每3秒生成一个令牌，逻辑是先计算给定的时间t减去上一次生成令牌的时间，如果t在last之前，则不会增加
	//新的令牌，等于说是毫无变化，维持原样，如果大于则会在这个时间差下按照原速率能产生多少令牌，如果多于令牌桶的
	//容量也会设置成令牌桶的容量，然后返回新的令牌桶，桶是满的，而速率是新的速率
	limiter.SetLimitAt(time.Now().Add(10*time.Second), rate.Every(3*time.Second))

	for i := 4; i < 9; i++ {
		log.Printf("got #%d,err:%v", i, limiter.Wait(context.Background()))
	}

}
