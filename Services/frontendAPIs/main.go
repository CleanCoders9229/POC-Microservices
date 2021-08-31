package main

import (
	"context"
	"flag"
	"log"
	"time"

	pb "github.com/CleanCoders9229/POC-Microservices/Services/proto/UserSchema"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

const (
	ginPortDefault  = ":3000"
	grpcAddrDefault = "localhost:50051"
)

func main() {
	// Flags Parser
	grpcAddr := flag.String("grpcAddr", grpcAddrDefault, "gRPC Address and Port.")
	ginPort := flag.String("ginPort", ginPortDefault, "Gin Server Port.")
	flag.Parse()

	// gRPC
	conn, err := grpc.Dial(*grpcAddr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("gRPC connection error: %v", err)
	}

	log.Printf("gRPC Open to: %v", *grpcAddr)
	defer conn.Close()
	manager := pb.NewRegistrationClient(conn)

	// GIN Server
	router := gin.Default()

	for i := 0; i < 2; i++ {
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
		time.Sleep(time.Second)
	}
}
