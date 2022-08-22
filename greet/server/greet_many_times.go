package main

import (
	"fmt"
	"log"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
)

// GreetManyTimes handles request from client and sends responses in a loop.
func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("GreetManyTimes function was invoked with: %v", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello, %s, number %d", in.FirstName, i)

		stream.Send(&pb.GreetResponse{
			Result: res,
		})
	}

	return nil
}
