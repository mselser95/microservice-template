package server

import (
	"context"
	pb "github.com/mselser95/microservice-template/internal/proto"
)

func (s *Server) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	// Example user response
	user := &pb.User{
		Id:    req.GetId(),
		Name:  "John Doe",
		Email: "johndoe@example.com",
	}
	return &pb.UserResponse{User: user}, nil
}
