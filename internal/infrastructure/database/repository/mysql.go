package repository

import (
	"ahuj/internal/domain/model"
	"ahuj/internal/domain/repository"
	"ahuj/internal/infrastructure/database/entity"
	"context"
	"database/sql"
	"fmt"
)

type MySQL struct {
	db *sql.DB
}

func (d *MySQL) OneByID(ctx context.Context, id int) (*model.Result, error) {
	fmt.Println("mysql OneByID")

	var m entity.Result

	row := d.db.QueryRowContext(ctx,
		"SELECT id, x, y, sum, created_at FROM results WHERE id = ?",
		id,
	)

	err := row.Scan(&m.ID, &m.X, &m.Y, &m.Sum, &m.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, fmt.Errorf("results one by id; scan error: %s", err)
	}

	return m.MapToDomain(), nil
}

func (d *MySQL) Save(ctx context.Context, result *model.Result) error {
	fmt.Println("mysql Save")

	_, err := d.db.ExecContext(
		ctx,
		"INSERT INTO results (x, y, sum) VALUES (?, ?, ?)",
		result.X, result.Y, result.Sum,
	)
	if err != nil {
		return fmt.Errorf("messages create; query error: %s", err)
	}

	return nil
}

func NewMySQLResultRepository(db *sql.DB) repository.Result {
	return &MySQL{db: db}
}
