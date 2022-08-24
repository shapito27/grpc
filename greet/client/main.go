package main

import (
	"log"

	pb "github.com/shapito27/grpc-go-course/greet/proto"
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

	c := pb.NewGreetServiceClient(conn)

	// grpc unary.
	//DoGreet(c)

	// grpc server streaming.
	//DoGreetManyTimes(c)

	// grpc client streaming.
	//DoLongGreet(c)

	// grpc Bi-directional streaming.
	DoGreetEvryone(c)
}
