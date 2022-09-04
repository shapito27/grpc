package main

import (
	"context"
	"fmt"
	"log"

	"github.com/golang/protobuf/ptypes/empty"
	pb "github.com/shapito27/grpc-go-course/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

// DeleteBlog handle request for delete blog.
func (s *Server) DeleteBlog(ctx context.Context, blogID *pb.BlogId) (*empty.Empty, error) {
	log.Printf("DeleteBlog was invoked with: %v\n", blogID)

	oid, err := primitive.ObjectIDFromHex(blogID.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can not parse ID")
	}

	res, err := collection.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Can not delete object from mongoDB with ID %v and error:%v", oid, err),
		)
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(
			codes.NotFound,
			"Blog was not found",
		)
	}

	return &emptypb.Empty{}, nil
}
