package main

import (
	"context"
	"io"
	"log"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
)

// DoGreetManyTimes sends request to server and receive responses from streaming server.
func DoGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("DoGreetManyTimes was invoked")
	req := &pb.GreetRequest{
		FirstName: "Ruslan",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)
	if err != nil {
		log.Fatalf("Failed GreetManyTimes, error: %v", err)
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Failed to receive from stream, error: %v", err)
		}

		log.Printf("GreetManyTimes: %s\n", resp)
	}
}
