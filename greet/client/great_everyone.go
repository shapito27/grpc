package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
)

// DoGreetEvryone sends request to server and receive responses from streaming server.
func DoGreetEvryone(c pb.GreetServiceClient) {
	log.Println("DoGreetEvryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Failed to greet everyone: %v", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Ruslan"},
		{FirstName: "Jane"},
		{FirstName: "Annet"},
		{FirstName: "Test"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request: %v", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			resp, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Failed to receive: %v\n", err)
			}

			log.Printf("Response is %s", resp.Result)
		}

		close(waitc)
	}()

	<-waitc
}
