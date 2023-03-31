package repository

import (
	"ahuj/internal/domain/model"
	"context"
)

type Result interface {
	OneByID(ctx context.Context, id int) (*model.Result, error)
	Save(ctx context.Context, result *model.Result) error
}
