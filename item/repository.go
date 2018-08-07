package item

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"
)

type Repository interface {
	Close()
	GetItemByID(ctx cotext.Context, id int32) (*Item, error)
	//ListItems(ctx context.Context, skip uint32, take uint32) ([]Item, error)
	//InsertItem(ctx cotext.Context, id int32) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil.err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) GetItemByID(ctx cotext.Context, id int32) (*Item, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, price FROM items WHERE id = $1", id)
	item := &Item{}
	if err := row.Scan(&a.ID, &a.Name); err != nil {
		return nil, err
	}
	return item, nil
}
