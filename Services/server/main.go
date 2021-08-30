package main

import (
	"context"
	"log"
	"net"

	pb "github.com/CleanCoders9229/POC-Microservices/Services/proto/UserSchema"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedRegistrationServer
}

func (s *server) CreateNewUser(ctx context.Context, in *pb.Profile) (*pb.Profile, error) {
	log.Printf("Receivee from client (CreateNewUser): %v.", in.Fullname)

	return &pb.Profile{Fullname: in.GetFullname(), Password: "", Email: in.GetEmail(), IsActivated: true, CreatedDate: true}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed at TCP:%v", port)
	}

	// TODO: Add gRPC Sever
	s := grpc.NewServer()
	pb.RegisterRegistrationServer(s, &server{})

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed at Serve: %v", err)
	}

	log.Println("===== START SERVER =====")
	log.Printf("running at: localhost%s", port)
}
