package main

import (
	"context"
	"log"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
)

// Greet handles requests.
func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet function was invoked with %v\n", in)
	return &pb.GreetResponse{
		Result: "Hello, " + in.FirstName,
	}, nil
}
