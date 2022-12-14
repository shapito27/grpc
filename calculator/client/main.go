package main

import (
	"log"

	pb "github.com/shapito27/grpc-go-course/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var addr string = "0.0.0.0:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial: %v\n", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorServiceClient(conn)

	//SendNums(c)

	//DoSqrt(c, 8)
	// check how client hadling eror.
	DoSqrt(c, -34)
}
