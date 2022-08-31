package main

import (
	"context"
	"log"

	pb "github.com/shapito27/grpc-go-course/blog/proto"
)

// update sends request for updating blog.
func update(c pb.BlogServiceClient, id string) {
	log.Println("update was invoked.")

	blog := &pb.Blog{
		Id:       id,
		AuthorId: "Ruslan",
		Title:    "new blog title",
		Content:  "new content",
	}

	_, err := c.UpdateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Can't update blog: %v\n", err)
	}

	log.Println("Blog was updated")
}
