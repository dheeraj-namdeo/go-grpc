package main

import (
	"fmt"
	"log"

	"github.com/dngithub/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am grpc client")

	clientConn, err := grpc.Dial("localhost:500051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial gRPC: %v", err)
	}
	defer clientConn.Close()
	c := greetpb.NewGreetServiceClient(clientConn)
	fmt.Printf("client created %v", c)
}
