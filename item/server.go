package item

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/masakazutakewaka/grpc-proto/item/pb"
)

type itemServer struct {
	service Service
}

func ListenGRPC(s Service, port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterItemServiceServer(server, &itemServer{s})
	reflection.Register(server)
	return server.Serve(listen)
}

func (s *itemServer) GetItem(ctx context.Context, r *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	item, err := s.r.GetItemByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetItemResponse{
		Item: &pb.Item{
			Id:    item.Id,
			Name:  item.Name,
			Price: item.Price,
		},
	}, nil
}

//func (s *itemServer) GetItems(ctx context.Context, r *pb.GetItemsRequest) (*pb.GetItemsResponse, error) { }

//func (s *itemServer) PostItem(ctx context.Context, r *pb.PostItemRequest) (*pb.PostItemResponse, error) { }
