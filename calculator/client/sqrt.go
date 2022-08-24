package main

import (
	"context"
	"log"

	pb "github.com/shapito27/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DoSqrt sends request to server.
func DoSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("DoSqrt was invoked!")

	resp, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Number: n,
	})

	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			log.Printf("Error message: %s, code: %s.", grpcErr.Message(), grpcErr.Code())
			if grpcErr.Code() == codes.InvalidArgument {
				log.Println("We probably sent negative number.")

			}
			return
		} else {
			log.Fatalf("Non gRPC error: %v\n", err)
		}
	}

	log.Printf("Result: %f\n", resp.Result)
}
