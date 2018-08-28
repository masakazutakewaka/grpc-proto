package coordinate

import (
	"fmt"
	"golang.org/x/net/context"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/masakazutakewaka/grpc-proto/coordinate/pb"
	"github.com/masakazutakewaka/grpc-proto/item"
	"github.com/masakazutakewaka/grpc-proto/user"
)

type coordinateServer struct {
	r          Repository
	itemClient *item.Client
	userClient *user.Client
}

func ListenGRPC(r Repository, itemURL string, userURL string, port int) error {
	itemClient, err := item.NewClient(itemURL)
	if err != nil {
		return err
	}

	userClient, err := user.NewClient(userURL)
	if err != nil {
		return err
	}

	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		itemClient.Close()
		userClient.Close()
		return err
	}

	server := grpc.NewServer()
	pb.RegisterCoordinateServiceServer(server, &coordinateServer{r, itemClient, userClient})
	reflection.Register(server)
	return server.Serve(listen)
}

func (s *coordinateServer) GetCoordinate(ctx context.Context, r *pb.GetCoordinateRequest) (*pb.GetCoordinateResponse, error) {
	coordinate, err := s.r.GetCoordinateByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetCoordinateResponse{Coordinate: coordinate}, nil
}

func (s *coordinateServer) GetCoordinatesByUser(ctx context.Context, r *pb.GetCoordinatesByUserRequest) (*pb.GetCoordinatesByUserResponse, error) {
	coordinates, err := s.r.GetCoordinatesByUserId(ctx, r.UserId)
	if err != nil {
		return nil, err
	}
	return &pb.GetCoordinatesResponse{Coordinates: coordinates}, nil
}

func (s *coordinateServer) PostCoordinate(ctx context.Context, r *pb.PostCoordinateRequest) (*pb.PostCoordinateResponse, error) {
	err := s.r.InsertCoordinate(ctx, r.UserId, r.ItemIds)
	if err != nil {
		return nil, err
	}
	return &pb.PostCoordinateResponse{}, nil
}