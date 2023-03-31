package repository

import "context"

type Adder interface {
	Do(ctx context.Context, x, y int) (int, error)
}
