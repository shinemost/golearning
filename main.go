package main

import (
	"context"
	"fmt"
	"log"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

var (
	maxWorkers = runtime.GOMAXPROCS(0)
	sema       = semaphore.NewWeighted(int64(maxWorkers))
	task       = make([]int, maxWorkers*4)
)

func main() {

	ctx := context.Background()
	for i := range task {
		if err := sema.Acquire(ctx, 1); err != nil {
			break
		}

		go func(i int) {
			defer sema.Release(1)
			time.Sleep(100 * time.Millisecond)
			task[i] = i + 1
		}(i)
	}

	err := sema.Acquire(ctx, int64(maxWorkers))
	if err != nil {
		log.Printf("获取所有的worker失败：%v", err)
	}
	fmt.Println(task)
}
