package main

import (
	"context"
	"log"

	pb "github.com/shapito27/grpc-go-course/calculator/proto"
)

// SendNums sends two numbers to server.
func SendNums(c pb.CalculatorServiceClient) {
	log.Println("SendNums was invoked")
	res, err := c.Calculator(context.Background(), &pb.CalculatorRequest{
		Num_1: 3,
		Num_2: 9,
	})
	if err != nil {
		log.Fatalf("Failed to SendNums, error: %v\n", err)
	}

	log.Printf("SendNums response: %s\v", res)
}
