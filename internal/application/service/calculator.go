package service

import (
	"ahuj/internal/domain/model"
	"ahuj/internal/domain/repository"
	"context"
	"fmt"
)

type Calculator struct {
	AdderRepository  repository.Adder
	ResultRepository repository.Result
}

func (c *Calculator) Add(ctx context.Context, x, y int) (int, error) {
	sum, err := c.AdderRepository.Do(ctx, x, y)
	if err != nil {
		return 0, fmt.Errorf("add error: %s", err)
	}

	r := &model.Result{
		X:   x,
		Y:   y,
		Sum: sum,
	}

	err = c.ResultRepository.Save(ctx, r)
	if err != nil {
		return 0, fmt.Errorf("save result error: %s", err)
	}

	return sum, nil
}

func (c *Calculator) History(ctx context.Context, id int) (*model.Result, error) {
	m, err := c.ResultRepository.OneByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("one by id error: %s", err)
	}

	return m, nil
}
