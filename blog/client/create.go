package main

import (
	"context"
	"log"

	pb "github.com/shapito27/grpc-go-course/blog/proto"
)

// createBlog creates blog.
func createBlog(c pb.BlogServiceClient) string {
	log.Println("crateBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Author id example",
		Title:    "My first blog",
		Content:  "Content of the blog",
	}

	res, err := c.CreateBlog(context.Background(), blog)
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	log.Printf("Blog was created: %s", res.Id)

	return res.Id
}
