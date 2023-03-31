package service

import (
	"ahuj/internal/domain/model"
	"context"
)

type Calculator interface {
	Add(ctx context.Context, x, y int) (int, error)
	History(ctx context.Context, id int) (*model.Result, error)
}
