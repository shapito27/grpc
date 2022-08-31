package main

import (
	"context"
	"io"
	"log"

	pb "github.com/shapito27/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
)

// list lists blogs.
func list(c pb.BlogServiceClient) {
	log.Println("list was invoked")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling listblog: %v", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Errror while receive response from stream: %v", err)
		}

		log.Println(res)
	}

}
