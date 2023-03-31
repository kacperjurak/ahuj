package entity

import (
	"ahuj/internal/domain/model"
	"time"
)

type Result struct {
	ID        int
	X         int
	Y         int
	Sum       int
	CreatedAt time.Time
}

func (m *Result) MapToDomain() *model.Result {
	return &model.Result{
		X:   m.X,
		Y:   m.Y,
		Sum: m.Sum,
	}
}
