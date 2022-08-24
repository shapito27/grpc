package main

import (
	"io"
	"log"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
)

// GreetEveryone.
func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone function was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Failed to receive: %v\n", err)
		}

		err = stream.Send(&pb.GreetResponse{
			Result: "Hello, " + req.FirstName,
		})

		if err != nil {
			log.Fatalf("Failed to send: %v\n", err)
		}

	}
}
