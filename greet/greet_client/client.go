package main

import (
	"context"
	"fmt"
	"io"
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
	// doUnary(c)
	doServerStreaming(c)
}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Started to do Unary RPC...")
	req := &greetpb.GreetRequest{
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

func doServerStreaming(c greetpb.GreetServiceClient) {
	fmt.Println("Started to do Server Streaming RPC...")
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Parsa",
			LastName:  "Akbari",
		},
	}
	res, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes RPC: %v", err)
	}
	for {
		msg, err := res.Recv()
		if err == io.EOF {
			break // Streaming is finished.
		}
		if err != nil {
			log.Fatalf("Error while Server Streaming: %v", err)
		}
		fmt.Println(msg.Result)
	}
}
