package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/shapito27/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// CreateBlog creates new blog.
func (s *Server) CreateBlog(ctx context.Context, blog *pb.Blog) (*pb.BlogId, error) {
	log.Printf("CreateBlog was invoked with %v\n", blog)

	data := BlogItem{
		AuthorID: blog.AuthorId,
		Title:    blog.Title,
		Content:  blog.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal error: %v", err))
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Can't convert to OID: %v", err))
	}

	return &pb.BlogId{
		Id: oid.Hex(),
	}, nil
}
