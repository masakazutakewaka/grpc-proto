package coordinate

import (
	"database/sql"
	"golang.org/x/net/context"

	"github.com/golang/glog"
	"github.com/lib/pq"

	"github.com/masakazutakewaka/grpc-proto/coordinate/pb"
)

const MaxItemCnt = 10

type Repository interface {
	Close()
	GetCoordinatesByUserId(ctx context.Context, userId int32) ([]*pb.Coordinate, error)
	InsertCoordinate(ctx context.Context, userId int32, itemIds []int32) error
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

func (r *postgresRepository) GetCoordinatesByUserId(ctx context.Context, userId int32) ([]*pb.Coordinate, error) {
	rows, err := r.db.QueryContext(ctx, "SELECT c.id, ci.item_id FROM coordinates AS c INNER JOIN coordinate_items AS ci ON c.id = ci.coordinate_id WHERE c.user_id = $1", userId)
	if err != nil {
		return nil, err
	}

	coordinates := []*pb.Coordinate{}
	itemsForCoordinate := make([][]int32, MaxItemCnt)

	for rows.Next() {
		var (
			itemId       int
			coordinateId int
		)

		if err := rows.Scan(&coordinateId, &itemId); err != nil {
			return nil, err
		}

		itemsForCoordinate[coordinateId] = append(itemsForCoordinate[coordinateId], int32(itemId))
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	for _, items := range itemsForCoordinate {
		if len(items) == 0 {
			continue
		}
		coordinate := &pb.Coordinate{}
		coordinate.UserId = userId
		coordinate.ItemIds = items
		coordinates = append(coordinates, coordinate)
	}

	return coordinates, nil
}

func (r *postgresRepository) InsertCoordinate(ctx context.Context, userId int32, itemIds []int32) error {
	var coordinateId int
	row := r.db.QueryRowContext(ctx, "INSERT INTO coordinates(user_id) VALUES($1) RETURNING id", userId)
	if err := row.Scan(&coordinateId); err != nil {
		glog.Error(err)
	}

	tx, err := r.db.Begin()
	if err != nil {
		glog.Error(err)
	}

	stmt, err := tx.Prepare(pq.CopyIn("coordinate_items", "coordinate_id", "item_id"))
	if err != nil {
		glog.Error(err)
	}

	for _, itemId := range itemIds {
		_, err = stmt.Exec(coordinateId, itemId)
		if err != nil {
			glog.Error(err)
		}
	}

	_, err = stmt.Exec()
	if err != nil {
		glog.Error(err)
	}

	err = stmt.Close()
	if err != nil {
		glog.Error(err)
	}

	err = tx.Commit()
	if err != nil {
		glog.Error(err)
	}

	return nil
}
