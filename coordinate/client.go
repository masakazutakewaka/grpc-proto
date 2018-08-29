package coordinate

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/masakazutakewaka/grpc-proto/coordinate/pb"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.CoordinateServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewCoordinateServiceClient(conn)
	return &Client{conn, client}, nil
}

func (client *Client) Close() {
	client.conn.Close()
}

func (client *Client) GetCoordinatesByUser(ctx context.Context, user_id int32) ([]*pb.Coordinate, error) {
	res, err := client.service.GetCoordinates(ctx, &pb.GetCoordinatesRequest{UserId: user_id})
	if err != nil {
		return nil, err
	}
	return res.Coordinates, nil
}

func (client *Client) PostCoordinate(ctx context.Context, user_id int32, item_ids []int32) error {
	_, err := client.service.PostCoordinate(ctx, &pb.PostCoordinateRequest{UserId: user_id, ItemIds: item_ids})
	return err
}