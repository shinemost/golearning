package main

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/pieterclaerhout/go-waitgroup"
)

func main() {

	ctx := context.Background()

	wg, _ := waitgroup.NewErrorGroup(ctx, 5)

	wg.Add(
		func() error {
			return nil
		},
		func() error {
			return errors.New("an error occurred")
		},
	)

	if err := wg.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

}
