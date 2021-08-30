package main

import (
	"context"
	"log"
	"time"

	pb "github.com/CleanCoders9229/POC-Microservices/Services/proto/UserSchema"
	"google.golang.org/grpc"
)

const (
	ginPort       = ":3000"
	ServerAddress = "localhost:50051"
)

func main() {
	// gRPC
	conn, err := grpc.Dial(ServerAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("gRPC connection error: %v", err)
	}
	defer conn.Close()
	manager := pb.NewRegistrationClient(conn)

	for true {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		req := &pb.Profile{
			Fullname:    "Ada De Sions",
			Password:    "1234",
			Email:       "ada@adacode.com",
			IsActivated: false,
			CreatedDate: false,
		}

		res, err := manager.CreateNewUser(ctx, req)

		if err != nil {
			log.Fatalf("Response Error: %v", err)
		}

		log.Printf("Server response: %s", res.String())
	}

	time.Sleep(2 * time.Second)

}
