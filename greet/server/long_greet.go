package main

import (
	"fmt"
	"io"
	"log"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
)

// LongGreet handles requests.
func (s *Server) LongGreet(stream pb.GreetService_LongGreetServer) error {
	log.Println("LongGreet func was invoked")

	res := ""

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Failed to receive request: %v", err)
		}

		res += fmt.Sprintf("Hello, %s!\n", req.FirstName)
	}
}
