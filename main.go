package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/vardius/gollback"
)

func main() {
	rs, errs := gollback.Race(context.Background(), func(ctx context.Context) (interface{}, error) {
		time.Sleep(3 * time.Second)
		//return 1, nil
		return nil, errors.New("failed1")
	}, func(ctx context.Context) (interface{}, error) {
		return nil, errors.New("failed2")
	}, func(ctx context.Context) (interface{}, error) {
		//return 3,nil;
		return nil, errors.New("failed3")
	})
	fmt.Println(rs)
	fmt.Println(errs)

}
