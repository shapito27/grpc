package main

import (
	"log"
	"time"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
	"golang.org/x/net/context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// DoGreatWithDeadline sends request to server with timeout.
func DoGreatWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("DoGreatWithDeadline was invoked.")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	resp, err := c.GreetWithDeadline(ctx, &pb.GreetRequest{
		FirstName: "Ruslan",
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline exceeded!")
			} else {
				log.Printf("Unexpected gRPC Error %s, code: %s", e.Err(), e.Code())
			}
		} else {
			log.Fatalf("Non gRPC error: %v", err)

		}
		return
	}

	log.Printf("DoGreatWithDeadline: %s\n", resp.Result)
}
