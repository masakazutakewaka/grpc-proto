package item

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"github.com/masakazutakewaka/grpc-proto/item/pb"
)

type Client struct {
	conn    *grpc.ClientConn
	service pb.ItemServiceClient
}

func NewClient(url string) (*Client, error) {
	conn, err := grpc.Dial(url, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	client := pb.NewItemServiceClient(conn)
	return &Client{conn, client}, nil
}

func (client *Client) Close() {
	client.conn.Close()
}

func (client *Client) GetItem(ctx context.Context, id int32) (*pb.Item, error) {
	res, err := client.service.GetItem(ctx, &pb.GetItemRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return res.Item, nil
}

func (client *Client) GetItems(ctx context.Context, ids []int32) ([]*pb.Item, error) {
	res, err := client.service.GetItems(ctx, &pb.GetItemsRequest{Ids: ids})
	if err != nil {
		return nil, err
	}
	return res.Items, nil
}

func (client *Client) PostItem(ctx context.Context, name string, price int32) error {
	_, err := client.service.PostItem(ctx, &pb.PostItemRequest{Name: name, Price: price})
	if err != nil {
		return err
	}
	return nil
}
