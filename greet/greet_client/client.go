package main

import (
	"context"
	"fmt"
	"log"

	"github.com/parsaakbari1209/gRPC-Golang-Hello-World/greet/greetpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Started client...")
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect server: %v", err)
	}
	defer conn.Close()
	c := greetpb.NewGreetServiceClient(conn)
	doUnary(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Started to do Unary RPC...")
	req := &greetpb.GreetingRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Parsa",
			LastName:  "Akbari",
		},
	}
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed in Greet unary RPC: %v", err)
	}
	fmt.Println(res.GetResult())
}
