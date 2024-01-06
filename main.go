package main

import "fmt"

var (
	a = c + b
	b = f()
	c = f()
	d = 3
)

func main() {
	fmt.Println(a, b, c, d)
}

func f() int {
	d++
	return d
}
