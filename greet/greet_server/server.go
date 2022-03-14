package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/dngithub/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	firstname := req.GetGreeting().FirstName
	result := "Hello " + firstname
	response := &greetpb.GreetResponse{Result: result}

	return response, nil
}

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
