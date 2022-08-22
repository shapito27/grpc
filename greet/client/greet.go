package main

import (
	"context"
	"log"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
)

// DoGreet send request, get response and print both.
func DoGreet(c pb.GreetServiceClient) {
	log.Println("DoGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Ruslan",
	})
	if err != nil {
		log.Fatalf("Failed to Greet, error: %v\n", err)
	}

	log.Printf("Greet response: %s\v", res)
}
