package main

import (
	"fmt"

	cmp "github.com/orcaman/concurrent-map/v2"
)

func main() {

	// 创建一个新的 map.
	m := cmp.New[string]()

	// 设置变量m一个键为“foo”值为“bar”键值对
	m.Set("foo", "bar")

	// 从m中获取指定键值.
	bar, ok := m.Get("foo")

	fmt.Printf("%s,%t\n", bar, ok)

	// 删除键为“foo”的项
	m.Remove("foo")

	bar, ok = m.Get("foo")

	fmt.Printf("%s,%t\n", bar, ok)

}
