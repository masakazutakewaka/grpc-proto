package user

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/masakazutakewaka/grpc-proto/user/pb"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.UserServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	client := pb.NewUserServiceClient(conn)
	return &Client{conn, client}, nil
}

func (client *Client) Close() {
	client.conn.Close()
}

func (client *Client) GetUser(ctx context.Context, id int32) (*pb.User, error) {
	res, err := client.service.GetUser(ctx, &pb.GetUserRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:   res.User.Id,
		Name: res.User.Name,
	}, nil
}

func (client *Client) GetUsers(ctx context.Context, skip int32, take int32) ([]*pb.User, error) {
	res, err := client.service.GetUsers(ctx, &pb.GetUsersRequest{Skip: skip, Take: take})
	if err != nil {
		return nil, err
	}
	return res.Users, nil
}

func (client *Client) PostUser(ctx context.Context, name string) error {
	err := client.service.PostUser(ctx, &pb.PostUserRequest{Name: name})
	if err != nil {
		return err
	}
	return nil
}
