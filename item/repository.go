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
	ListItems(ctx context.Context, skip uint32, take uint32) ([]pb.Item, error)
	InsertItem(ctx context.Context, name string, price int32) error
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

func (r *postgresRepository) ListItems(ctx context.Context, skip uint32, take uint32) ([]pb.Item, error) {
	rows := r.db.QueryContext(ctx, "SELECT id, name, price FROM items OFFSET $1 LIMIT $2", skip, take)
	items := []pb.Item{}
	for rows.Next {
		item = pb.Item{}
		err = rows.Scan(&item.id, &item.name, &item.price)
		items = append(items, item)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *postgresRepository) InsertItem(ctx context.Context, name string, price int32) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO items(id, name, price) VALUES($1, $2)", name, price)
	return err
}