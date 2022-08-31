package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/shapito27/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// ListBlog returns list of blogs.
func (s *Server) ListBlog(in *empty.Empty, stream pb.BlogService_ListBlogServer) error {
	log.Println("listblog was invoked.")

	cur, err := collection.Find(context.Background(), primitive.D{{}})
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Unknown internal error: %v", err))
	}
	defer cur.Close(context.Background())

	for cur.Next(context.Background()) {
		data := &BlogItem{}
		err := cur.Decode(data)
		if err != nil {
			return status.Errorf(codes.Internal, fmt.Sprintf("Error while decoding data from MongoDB: %v", err))
		}

		stream.Send(documentToBlog(data))
	}

	if err = cur.Err(); err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error while getting next element from Cursor: %v", err))
	}

	return nil
}
