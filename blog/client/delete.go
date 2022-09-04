package main

import (
	"context"
	"log"

	pb "github.com/shapito27/grpc-go-course/blog/proto"
)

// deleteBlog deletes blog.
func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("deleteBlog was invoked!")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("error while deleting: %v\n", err)
	}

	log.Println("Blog was deleted!")
}
