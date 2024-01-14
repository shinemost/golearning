package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aaronjan/hunch"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := hunch.Waterfall(ctx, func(ctx context.Context, i interface{}) (interface{}, error) {
		n := 1
		return n * 2, nil
		//return nil, errors.New("failed")
	}, func(ctx context.Context, i interface{}) (interface{}, error) {
		fmt.Println("perivous fun result is : ", i)
		//return nil, errors.New("failed")
		return 4, nil
	})

	fmt.Println(res)
	fmt.Println(err)
}
