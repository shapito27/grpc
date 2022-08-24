package main

import (
	"context"
	"fmt"
	"log"
	"math"

	pb "github.com/shapito27/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Sqrt handles request and try to sqrt given number.
func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Printf("Sqrt was invoked with: %v\n", in)

	number := in.Number

	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Can't sqrt on negative: %d", number))
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
