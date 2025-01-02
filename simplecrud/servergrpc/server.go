package main

import (
	"context"
	"net"
	"sync"

	pb "femas66/simple-grpc/simplecrud"

	grpc "google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedCRUDServiceServer
	data []*pb.User
	mu   sync.Mutex
}

func (s *Server) CreateUser(ctx context.Context, req *pb.User) (*pb.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = append(s.data, req)
	return &pb.Response{Message: "Created Successfully"}, nil
}

func (s *Server) ReadUsers(ctx context.Context, req *pb.Empty) (*pb.UserList, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	return &pb.UserList{List: s.data}, nil
}

func (s *Server) UpdateUser(ctx context.Context, req *pb.User) (*pb.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, u := range s.data {
		if u.Id == req.Id {
			s.data[i] = req
			return &pb.Response{Message: "Updated Successfully"}, nil
		}
	}
	return &pb.Response{Message: "User Not Found"}, nil
}

func (s *Server) DeleteUser(ctx context.Context, req *pb.UserID) (*pb.Response, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	for i, u := range s.data {
		if u.Id == req.Id {
			s.data = append(s.data[:i], s.data[i+1:]...)
			return &pb.Response{Message: "Deleted Successfully"}, nil
		}
	}
	return &pb.Response{Message: "User Not Found"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	pb.RegisterCRUDServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		panic(err)
	}

}
