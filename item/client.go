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
	conn, err := grpc.Dial(url)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	client := pb.NewItemServiceClient(conn)
	return &Client{conn, client}, nil
}

func (client *Client) GetItem(ctx Context.context, id int32) (*pb.Item, error) {
	res, err := client.GetItem(ctx, &pb.GetItemRequest{Id: id})
	if err != nil {
		return nil, err
	}
	return &pb.Item{
		Id:    res.Item.Id,
		Name:  res.Item.Name,
		Price: res.Item.Price,
	}, nil
}

func (client *Client) GetItems(ctx Context.context, skip int32, take int32) (*pb.Items, error) {
	res, err := client.GetItem(ctx, &pb.GetItemRequest{Skip: skip, Take: take})
	if err != nil {
		return nil, err
	}
	return &pb.Items{
		Items: res.Items,
	}, nil
}

func (client *Client) PostItem(ctx Context.context, name string, price int32) (*pb.Item, error) {
	res, err := client.PostItem(ctx, &pb.PostItemRequest{Name: name, Price: price})
	if err != nil {
		return nil, err
	}
	return &pb.Item{
		Id:    res.Item.Id,
		Name:  res.Item.Name,
		Price: res.Item.Price,
	}, nil
}
