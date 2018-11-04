package user

import (
	"fmt"
	"golang.org/x/net/context"
	"net"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/masakazutakewaka/grpc-proto/user/pb"
)

type userServer struct {
	r Repository
}

func ListenGRPC(r Repository, port int) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return err
	}
	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userServer{r})
	reflection.Register(server)
	return server.Serve(listen)
}

func (s *userServer) GetUser(ctx context.Context, r *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, err := s.r.GetUserByID(ctx, r.Id)
	if err != nil {
		return nil, err
	}
	return &pb.GetUserResponse{User: user}, nil
}

func (s *userServer) GetUsers(ctx context.Context, r *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	users, err := s.r.ListUsers(ctx, r.Skip, r.Take)
	if err != nil {
		return nil, err
	}
	return &pb.GetUsersResponse{Users: users}, nil
}

func (s *userServer) PostUser(ctx context.Context, r *pb.PostUserRequest) (*empty.Empty, error) {
	err := s.r.InsertUser(ctx, r.Name)
	if err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}
