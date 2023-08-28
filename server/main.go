package main

import (
	pb "Protos/go/proto"
	"context"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	pb.UnimplementedTestServiceServer
}

func (s *Server) Send(ctx context.Context, point *pb.Point) (*pb.Response, error) {
	return &pb.Response{Name: fmt.Sprint("Response: ", point.Name)}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println("Error initializing server")
	}
	serv := &Server{}

	grpc := grpc.NewServer()
	pb.RegisterTestServiceServer(grpc, serv)

	if err := grpc.Serve(lis); err != nil {
		fmt.Println("Error starting server")
	}
	fmt.Println("Server running: 8000")

}
