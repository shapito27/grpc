package main

import (
	"context"
	"log"
	"time"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// GreetWithDeadline handles requests.
func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("GreetWithDeadline was invoked with: %s", in.FirstName)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			message := "The client canceled request."
			log.Println(message)
			return nil, status.Error(codes.Canceled, message)
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Result: "Hello, " + in.FirstName,
	}, nil
}
