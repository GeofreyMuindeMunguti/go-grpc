package main

import (
	pb "Protos/go/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, _ := grpc.Dial("localhost:8000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	client := pb.NewTestServiceClient(conn)
	response, err := client.Send(context.Background(), &pb.Point{ID: 1, Name: "MM"})
	if err != nil {
		fmt.Println("Error: ", err)
	}
	fmt.Println(response)
}
