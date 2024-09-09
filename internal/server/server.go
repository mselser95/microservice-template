package server

import (
	"go.uber.org/zap"
	"net"

	"github.com/mselser95/microservice-template/internal/config"
	pb "github.com/mselser95/microservice-template/internal/proto"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedUserServiceServer
}

func NewServer() *Server {
	return &Server{}
}

func StartServer(
	cfg config.Config,
	server *Server,
	logger *zap.Logger,
) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		logger.Error("failed-to-listen", zap.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, server)
	if err := s.Serve(lis); err != nil {
		logger.Error("failed-to-serve", zap.Error(err))
	}
}
