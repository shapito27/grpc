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

// UpdateBlog handles requests for update blog.
func (s *Server) UpdateBlog(ctx context.Context, blog *pb.Blog) (*empty.Empty, error){
	log.Printf("UpdateBlog was invoked with: %v\n", blog)

	oid, err := primitive.ObjectIDFromHex(blog.Id)
	if err!=nil {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Cant parse ID"))
	}

	data := &BlogItem{
		AuthorID: blog.AuthorId,
		Title: blog.Title,
		Content: blog.Content,
	}
	filter := bson.M{"_id": oid}
	res, err:=collection.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err!=nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cant update blog"))
	}
	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, fmt.Sprintf("Blog wasn't find"))
	}

	return &emptypb.Empty{}, nil
}
