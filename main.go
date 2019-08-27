package main

import (
	"context"
	"log"
	"net"

	"github.com/mahendrakariya/add/add"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

// server is used to implement helloworld.GreeterServer.
type server struct{}

func (s *server) DoSum(ctx context.Context, in *add.Numbers) (*add.Resp, error) {
	return &add.Resp{
		Sum: in.A + in.B,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	add.RegisterAdderServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
