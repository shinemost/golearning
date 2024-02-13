package main

import (
	"log"
	"time"
)

func main() {
	AfterFuncDemo()
}

// WaitChannel 设置超时时间
func WaitChannel(conn <-chan string) bool {
	timer := time.NewTimer(time.Second)

	select {
	case <-conn:
		timer.Stop()
		return true
	case <-timer.C:
		println("waitChannel timeout!")
		return false
	}
}

// DelayDoSomething 延迟执行逻辑
func DelayDoSomething() {
	timer := time.NewTimer(5 * time.Second)

	select {
	case <-timer.C:
		println("delay 5 seconds to do something!")

	}
}

// AfterDemo 简单接口After，等待对应的时间
func AfterDemo() {
	log.Println(time.Now())
	<-time.After(time.Second)
	log.Println(time.Now())
}

// AfterFuncDemo 延迟执行函数
func AfterFuncDemo() {
	log.Println(time.Now())
	_ = time.AfterFunc(2*time.Second, func() {
		log.Println("after func end:", time.Now())
	})

	//可以取消定时器
	//timer.Stop()
	time.Sleep(2 * time.Second)

}
