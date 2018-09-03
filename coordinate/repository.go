package coordinate

import (
	"database/sql"
	"golang.org/x/net/context"
	"strconv"
	"strings"

	_ "github.com/lib/pq"

	"github.com/masakazutakewaka/grpc-proto/coordinate/pb"
)

type Repository interface {
	Close()
	GetCoordinatesByUserId(ctx context.Context, user_id int32) ([]*pb.Coordinate, error)
	InsertCoordinate(ctx context.Context, user_id int32, item_ids []int32) error
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

func (r *postgresRepository) GetCoordinatesByUserId(ctx context.Context, user_id int32) ([]*pb.Coordinate, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT id, item_ids FROM coordinates WHERE user_id = $1", user_id)
	coordinates := []*pb.Coordinate{}
	if err != nil {
		return nil, err
	}

	var item_ids string

	for rows.Next() {
		coordinate := &pb.Coordinate{}
		if err := rows.Scan(&coordinate.Id, &item_ids); err != nil {
			return nil, err
		}
		coordinate.UserId = user_id
		coordinate.ItemIds, err = SliceItemIds(item_ids)
		if err != nil {
			return nil, err
		}
		coordinates = append(coordinates, coordinate)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return coordinates, nil
}

func (r *postgresRepository) InsertCoordinate(ctx context.Context, user_id int32, item_ids []int32) error {
	_, err := r.db.ExecContext(ctx, "INSERT INTO coordinates(user_id, item_ids) VALUES($1, $2)", user_id, BundleItemIds(item_ids))
	return err
}

func SliceItemIds(item_ids string) ([]int32, error) {
	sliced := []int32{}
	splited := strings.Split(item_ids, ", ")
	for _, item_id := range splited {
		strId, err := strconv.Atoi(item_id)
		if err != nil {
			return nil, err
		}
		sliced = append(sliced, int32(strId))
	}
	return sliced, nil
}

func BundleItemIds(item_ids []int32) string {
	bundled := []string{}
	for _, item_id := range item_ids {
		bundled = append(bundled, strconv.Itoa(int(item_id)))
	}
	return strings.Join(bundled, ", ")
}
