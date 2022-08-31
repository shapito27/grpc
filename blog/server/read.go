package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/shapito27/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ReadBlog handle requests for reading blogs.
func (s *Server) ReadBlog(ctx context.Context, blogID *pb.BlogId) (*pb.Blog, error) {
	log.Printf("ReadBlog was invoked with blogId %v", blogID)

	oid, err := primitive.ObjectIDFromHex(blogID.Id)
	if err != nil {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Can't parse id: %v", blogID),
		)
	}

	data := &BlogItem{}
	filter := bson.M{"_id": oid}
	res := collection.FindOne(ctx, filter)

	if err := res.Decode(data); err != nil {
		return nil, status.Errorf(
			codes.NotFound,
			fmt.Sprintf("Blog not found by id: %v", blogID),
		)
	}

	return documentToBlog(data), nil
}
