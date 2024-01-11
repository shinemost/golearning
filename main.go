package main

import (
	"context"
	"fmt"
	"sync/atomic"
	"time"

	"github.com/go-pkgz/syncs"
)

func main() {
	//默认是控制子任务的并发数量
	swg := syncs.NewSizedGroup(10)
	// 另一种处理方式，控制同时处理任务的协程数量
	//swg := syncs.NewSizedGroup(10, syncs.Preemptive)
	var c uint32

	for i := 0; i < 1000; i++ {
		swg.Go(func(ctx context.Context) {
			time.Sleep(5 * time.Millisecond)
			atomic.AddUint32(&c, 1)
		})
	}
	swg.Wait()
	fmt.Println(c)

}
