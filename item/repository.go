package item

import (
	"context"
	"database/sql"

	_ "github.com/lib/pq"

	"github.com/masakazutakewaka/grpc-proto/item/pb"
)

type Repository interface {
	Close()
	GetItemByID(ctx context.Context, id int32) (*pb.Item, error)
	//ListItems(ctx context.Context, skip uint32, take uint32) ([]Item, error)
	//InsertItem(ctx context.Context, id int32) error
}

type postgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(url string) (Repository, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &postgresRepository{db}, nil
}

func (r *postgresRepository) Close() {
	r.db.Close()
}

func (r *postgresRepository) Ping() error {
	return r.db.Ping()
}

func (r *postgresRepository) GetItemByID(ctx context.Context, id int32) (*pb.Item, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name, price FROM items WHERE id = $1", id)
	item := &pb.Item{}
	if err := row.Scan(&item.Id, &item.Name); err != nil {
		return nil, err
	}
	return item, nil
}
