package main

import (
	"fmt"
	"log"
	"net"

	"github.com/dngithub/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func main() {
	fmt.Println("Hello world")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen grpc: %v", err)
	}
	s := grpc.NewServer()
	//greetpb.RegisterGreetServiceServer
	greetpb.RegisterGreetServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve grpc: %v", err)
	}

}
