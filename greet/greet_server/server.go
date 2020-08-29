package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/parsaakbari1209/gRPC-Golang-Hello-World/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetingRequest) (*greetpb.GreetingResponse, error) {
	firstName := req.GetGreeting().GetFirstName()
	// lastName := req.GetGreeting().GetLastName()
	result := "Hello " + firstName
	res := &greetpb.GreetingResponse{
		Result: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Started server...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
