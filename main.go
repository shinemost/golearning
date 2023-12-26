package main

import (
	"fmt"

	"github.com/oxtoacart/bpool"
)

var bufpool *bpool.BufferPool

func main() {

	//数据库连接池和http client连接池底层都是用了Sync.Pool技术，有时间可以看看源码学习一下
	//sql.DB{}
	//http.Transport{}

	//三方包
	//bytebufferpool 底层是sync.Pool https://github.com/valyala/bytebufferpool/
	//oxtoacart/bpool 底层基于channel实现的 https://github.com/oxtoacart/bpool

	bufpool = bpool.NewBufferPool(48)

	fmt.Println(bufpool.NumPooled())

	buf := bufpool.Get()

	fmt.Println(bufpool.NumPooled())

	bufpool.Put(buf)

	fmt.Println(bufpool.NumPooled())
}

func someFunction() error {

	// Get a buffer from the pool
	buf := bufpool.Get()
	//...
	//...
	//...
	// Return the buffer to the pool
	bufpool.Put(buf)

	return nil
}
