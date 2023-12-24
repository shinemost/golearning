package main

import (
	"fmt"
	cmp "github.com/cornelk/hashmap"
)

func main() {

	m := cmp.New[uint8, int]()
	m.Set(1, 123)
	value, ok := m.Get(1)

	fmt.Printf("value is: %d,ok is %t\n", value, ok)

	n := cmp.New[string, int]()
	n.Set("amount", 456)
	value, ok = n.Get("amount")

	fmt.Printf("value is: %d,ok is %t\n", value, ok)

}
