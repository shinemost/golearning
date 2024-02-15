package main

import "fmt"

func main() {
	demo("hello")
	demo("hello", "leo", "yinruoliang")
	demo("hello", "s", 12, 32)
	g := []any{"caicai", "jjbong", "ruirui", 1, 2}
	demo("not happy", g...)
}

func demo(l string, who ...any) {
	if who == nil {
		fmt.Println("nobody say hi")
	}
	for _, s := range who {
		fmt.Printf("%s %v\n", l, s)
	}

}
