package main

import (
	"context"
	"log"

	pb "github.com/shapito27/grpc-go-course/calculator/proto"
)

// Calculator handles request from client, return sum of two given numbers.
func (s *Server) Calculator(ctx context.Context, in *pb.CalculatorRequest) (*pb.CalculatorResponse, error) {
	log.Printf("Calculator was invoked with %v", in)
	return &pb.CalculatorResponse{
		Sum: in.Num_1 + in.Num_2,
	}, nil
}
