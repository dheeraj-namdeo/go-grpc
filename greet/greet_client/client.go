package main

import (
	"context"
	"fmt"
	"log"

	"github.com/dngithub/go-grpc/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello i am grpc client")

	clientConn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Unable to dial gRPC: %v", err)
	}
	defer clientConn.Close()
	c := greetpb.NewGreetServiceClient(clientConn)
	//fmt.Printf("client created %v", c)

	doUnaryrpc(c)

}

func doUnaryrpc(c greetpb.GreetServiceClient) {
	fmt.Println("starting doUnaryRpc client")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{FirstName: "dheeraj", LastName: "namdeo"},
	}
	response, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error in GreetResponse: %v", err)
	}
	fmt.Printf("Result of greetResponse: %v", response.Result)
}
