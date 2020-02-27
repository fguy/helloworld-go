package helloworld

import (
	"context"
	"database/sql"

	"github.com/fguy/helloworld-go/entities"
	"go.uber.org/fx"
)

const (
	query = `
		SELECT 
			body
		FROM pages
		WHERE title = ?`
)

type repository struct {
	db *sql.DB
}

func (r *repository) GetPage(ctx context.Context, title string) (*entities.Page, error) {
	result := &entities.Page{
		Title: title,
	}

	row := r.db.QueryRowContext(ctx, query, title)

	if err := row.Scan(
		&result.Body,
	); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return result, nil
}

// New -
func New(lc fx.Lifecycle, openDB func() (*sql.DB, error)) (Interface, error) {
	db, err := openDB()
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(context.Context) error {
			return db.Close()
		},
	})

	return &repository{
		db: db,
	}, nil
}
