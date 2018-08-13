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
	r Repository
}

func ListenGRPC(r Repository, port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterItemServiceServer(server, &itemServer{r})
	reflection.Register(server)
	return server.Serve(listen)
}

func (s *itemServer) GetItem(ctx context.Context, r *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	item, err := s.r.GetItemByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetItemResponse{Item: item}, nil
}

func (s *itemServer) GetItems(ctx context.Context, r *pb.GetItemsRequest) (*pb.GetItemsResponse, error) {
	items, err := s.r.ListItems(ctx, r.skip, r.take)
	if err != nil {
		return nil, err
	}
	return &pb.GetItemsResponse{Item: items}, nil
}

func (s *itemServer) PostItem(ctx context.Context, r *pb.PostItemRequest) (*pb.PostItemResponse, error) {
	item, err := s.r.InsertItem(ctx, r.name, r.price)
	if err != nil {
		return nil, err
	}
	return &pb.PostItemResponse{Item: item}, nil
}
