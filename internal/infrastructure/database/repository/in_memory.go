package repository

import (
	"ahuj/internal/domain/model"
	"ahuj/internal/infrastructure/database/entity"
	"context"
	"fmt"
	"time"
)

type InMemory struct {
	results map[int]entity.Result
}

func (d *InMemory) OneByID(_ context.Context, id int) (*model.Result, error) {
	fmt.Println("in memory OneByID")

	r, ok := d.results[id]
	if !ok {
		return nil, fmt.Errorf("no record with id: %d", id)
	}

	return &model.Result{
		X:   r.X,
		Y:   r.Y,
		Sum: r.Sum,
	}, nil
}

func (d *InMemory) Save(_ context.Context, result *model.Result) error {
	fmt.Println("in memory Save")

	l := len(d.results) + 1 // to start ids from 1, not 0
	d.results[l] = entity.Result{
		ID:        l,
		X:         result.X,
		Y:         result.Y,
		Sum:       result.Sum,
		CreatedAt: time.Now(), // yes, I know, it should be mocked
	}

	return nil
}

func NewInMemory() *InMemory {
	r := map[int]entity.Result{}
	return &InMemory{results: r}
}
