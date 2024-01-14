package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/aaronjan/hunch"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := hunch.All(ctx, func(ctx context.Context) (interface{}, error) {
		time.Sleep(3 * time.Second)
		return 1, nil
	}, func(ctx context.Context) (interface{}, error) {
		//return 2, nil
		return nil, errors.New("failed")
	}, func(ctx context.Context) (interface{}, error) {
		return 3, nil
	})

	fmt.Println(res)
	fmt.Println(err)
}
