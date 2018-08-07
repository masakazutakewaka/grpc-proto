package item

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"wear-proto/item/pb"
)

type itemServer struct {
	r Repository
}

func ListenGRPC(r Repository, port int) error {
	listen, err := net.listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterItemServer(server, &itemServer{r})
	reflection.Register(server)
	return server.Serve(listen)
}

func (s *itemServer) GetItem(ctx context.Context, r *pb.GetItemRequest) (*pb.GetItemResponse, error) {
	item, err := s.r.GetItemByID(r.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &pb.GetItemResponse{
		Id:    item.ID,
		Name:  item.Name,
		Price: item.Price,
	}
}

//func (s *itemServer) GetItems(ctx context.Context, r *pb.GetItemsRequest) (*pb.GetItemsResponse, error) { }

//func (s *itemServer) PostItem(ctx context.Context, r *pb.PostItemRequest) (*pb.PostItemResponse, error) { }
