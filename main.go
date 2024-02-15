package main

import (
	"fmt"
)

func main() {
	demo()
}

func demo() {
	i, j := 0, 0
	if true {
		j, k := 1, 1
		fmt.Println(j, k)
	}
	fmt.Println(i, j)
}
