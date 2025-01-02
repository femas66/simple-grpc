package main

import (
	"context"
	pb "femas66/simple-grpc/login"
	"log"
	"net"

	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedLoginServiceServer
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	if req.Username == "admin" && req.Password == "123456" {
		return &pb.LoginResponse{
			Success: true,
			Message: "Login Successful",
		}, nil
	}

	return &pb.LoginResponse{
		Success: false,
		Message: "Invalid Credentials",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterLoginServiceServer(s, &Server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
