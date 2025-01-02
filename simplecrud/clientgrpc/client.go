package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "femas66/simple-grpc/simplecrud"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewCRUDServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = client.CreateUser(ctx, &pb.User{Id: "1", Name: "Item 1"})
	if err != nil {
		log.Fatalf("Failed to create item: %v", err)
	}
	fmt.Println("Item created successfully")

	items, err := client.ReadUsers(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to read items: %v", err)
	}
	fmt.Println("Items:", items)

	_, err = client.UpdateUser(ctx, &pb.User{Id: "1", Name: "Updated Item 1"})
	if err != nil {
		log.Fatalf("Failed to update item: %v", err)
	}
	fmt.Println("Item updated successfully")

	_, err = client.DeleteUser(ctx, &pb.UserID{Id: "1"})
	if err != nil {
		log.Fatalf("Failed to delete item: %v", err)
	}
	fmt.Println("Item deleted successfully")
}
