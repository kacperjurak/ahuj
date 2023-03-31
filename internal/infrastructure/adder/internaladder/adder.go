package internaladder

import (
	"context"
	"fmt"
)

type Adder struct{}

func (a Adder) Do(_ context.Context, x, y int) (int, error) {
	fmt.Println("internal adder")

	return x + y, nil
}
