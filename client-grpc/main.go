package main

import (
	"context"
	"log"

	"google.golang.org/grpc"

	pb "femas66/simple-grpc/login"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewLoginServiceClient(conn)
	ctx := context.Background()
	req := &pb.LoginRequest{
		Username: "adm",
		Password: "1234",
	}
	resp, err := c.Login(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Login Response: %v", resp)
}
