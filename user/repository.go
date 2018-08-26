package user

import (
	"database/sql"
	"golang.org/x/net/context"

	_ "github.com/lib/pq"

	"github.com/masakazutakewaka/grpc-proto/user/pb"
)

type Repository interface {
	Close()
	GetUserByID(ctx context.Context, id int32) (*pb.User, error)
	ListUsers(ctx context.Context, skip int32, take int32) ([]*pb.User, error)
	InsertUser(ctx context.Context, name string) error
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

func (r *postgresRepository) GetUserByID(ctx context.Context, id int32) (*pb.User, error) {
	row := r.db.QueryRowContext(ctx, "SELECT id, name FROM users WHERE id = $1", id)
	user := &pb.User{}
	if err := row.Scan(&user.Id, &user.Name); err != nil {
		return nil, err
	}
	return user, nil
}

func (r *postgresRepository) ListUsers(ctx context.Context, skip int32, take int32) ([]*pb.User, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, name FROM users OFFSET $1 LIMIT $2", skip, take)
	users := []*pb.User{}
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		user := &pb.User{}
		if err := rows.Scan(&user.Id, &user.Name); err != nil {
			break
		}
		users = append(users, user)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func (r *postgresRepository) InsertUser(ctx context.Context, name string) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO users(name) VALUES($1)", name)
	return err
}
