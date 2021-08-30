package main

import (
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed at TCP:%v", port)
	}

	// TODO: Add gRPC Sever
}
