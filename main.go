package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"slices"
	"sync"
)

func main() {
	newHttpRoute()
}

// 1.22 之前range遍历只会创建一次变量，后面循环都是更新
// 1.22 每次都创建新变量
func newRange() {
	wg := sync.WaitGroup{}
	v := []string{"a", "b", "c"}
	wg.Add(len(v))
	for _, i := range v {
		// i每次都是新内存地址
		fmt.Printf("%p\n", &i)
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()

	// 新的遍历方式 从【0,3）从0到2
	for l := range 3 {
		fmt.Println(l)
	}
}

func newSliceMethod() {
	m := []int{1, 2, 3}
	n := []int{4, 5, 6}
	l := []int{7, 8, 9}

	// 1.22 之前多切片只能通过两两append
	merge := append(m, n...)
	last := append(merge, l...)
	fmt.Println(last)

	// 1.22 通过新增的切片方法，一行代码支持多个切片拼接
	// 内部其实也是通过遍历 + append实现的，封装好了直接拿来用，且支持泛型
	s := slices.Concat(m, n, l)

	// 切片反转
	slices.Reverse(s)
	fmt.Println(s)

	// 切片去重
	o := []int{1, 1, 2, 2, 3, 4, 4, 4, 5}
	no := slices.Compact(o)
	fmt.Println(no)
}

// http 路由新变化，支持restful，且可指定METHOD
func newHttpRoute() {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /hello/{name}", func(writer http.ResponseWriter, request *http.Request) {
		//if request.Method != "GET" {
		//	fmt.Fprint(writer, "warn: 只支持GET方法")
		//} else {
		fmt.Fprint(writer, "你好"+request.PathValue("name"))
		//}
	})

	if err := http.ListenAndServe(":5678", mux); err != nil {
		panic(err)
	}

}
