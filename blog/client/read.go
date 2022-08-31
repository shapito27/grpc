package main

import (
	"context"
	"log"

	pb "github.com/shapito27/grpc-go-course/blog/proto"
)

// readBlog reads blog.
func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("ReadBlog was invoked.")

	res, err := c.ReadBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error while read blog: %v", err)
	}

	log.Printf("Blog was read: %v\n", res)

	return res
}
