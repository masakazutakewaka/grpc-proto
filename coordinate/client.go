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

func (client *Client) GetCoordinatesByUser(ctx context.Context, userId int32) ([]*pb.Coordinate, error) {
	res, err := client.service.GetCoordinatesByUser(ctx, &pb.GetCoordinatesByUserRequest{UserId: userId})
	if err != nil {
		return nil, err
	}
	return res.Coordinates, nil
}

func (client *Client) PostCoordinate(ctx context.Context, userId int32, itemIds []int32) error {
	_, err := client.service.PostCoordinate(ctx, &pb.PostCoordinateRequest{UserId: userId, ItemIds: itemIds})
	if err != nil {
		return err
	}
	return nil
}
