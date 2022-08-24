package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
)

// DoLongGreet sends streaming requests, get response and print both.
func DoLongGreet(c pb.GreetServiceClient) {
	log.Println("DoLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Ruslan"},
		{FirstName: "Test"},
		{FirstName: "Andrei"},
		{FirstName: "Jacob"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("Failed to send LongGreet: %v", err)

	}

	for _, req := range reqs {
		log.Printf("Sending: %v", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to Close and recv: %v", err)

	}

	log.Printf("LongGreet res: %s\n", res.Result)

}
