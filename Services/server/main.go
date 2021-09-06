package main

import (
	"context"
	"log"
	"net"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	pb "github.com/CleanCoders9229/POC-Microservices/Services/proto/UserSchema"
	"google.golang.org/api/option"
	"google.golang.org/grpc"
)

const (
	PORT       = ":50051"
	TOKEN_PATH = "/Users/adax/go/src/Laboratory/POC-Microservices/Services/server/token.json"
)

type server struct {
	pb.UnimplementedRegistrationServer
}

var opt = option.WithCredentialsFile(TOKEN_PATH)
var app *firebase.App

func (s *server) CreateNewUser(ctx context.Context, in *pb.Profile) (*pb.Profile, error) {
	log.Printf("Received from client (CreateNewUser): %v.", in.String())

	client, err := app.Auth(ctx)
	if err != nil {
		log.Fatalf("error getting Auth client: %v\n", err)
	}

	params := (&auth.UserToCreate{}).
		Email(in.GetEmail()).
		EmailVerified(false).
		Password(in.GetPassword()).
		DisplayName(in.GetFullname()).
		PhotoURL("http://www.example.com/12345678/photo.png").
		Disabled(false)

	userProfile := &pb.Profile{Fullname: in.GetFullname(), Password: "****", Email: in.GetEmail(), IsActivated: true, CreatedDate: true}

	_, err = client.CreateUser(ctx, params)
	if err != nil {
		log.Printf("Response from creating user: %v\n", err)
		log.Printf("==========")
		log.Println(err)
		return nil, err
	}
	log.Printf("Successfully created user: %v\n", userProfile.String())

	return userProfile, nil
}

func main() {
	listen, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalf("Failed at TCP:%v", PORT)
	}

	// Start gRPC Sever
	s := grpc.NewServer()
	pb.RegisterRegistrationServer(s, &server{})

	// Connect to Firebase
	app, err = firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed at Serve: %v", err)
	}

	log.Println("===== START SERVER =====")
	log.Printf("running at: localhost%s", PORT)
}
