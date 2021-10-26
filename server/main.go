package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/toffernator/chitty-chat/client/bindings"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type Server struct {
	clients []string
	bindings.UnimplementedClientToServerServiceServer
}

func (s *Server) Join(ctx context.Context, in *bindings.Address) (*bindings.StatusOk, error) {
	fmt.Printf("Client %s joining server %s", in.Address, address)
	return &bindings.StatusOk{
		LamportTs: 0,
	}, nil
}

func (s *Server) Leave(ctx context.Context, in *bindings.Address) (*bindings.StatusOk, error) {
	fmt.Printf("Client %s leaving server %s", in.Address, address)
	return &bindings.StatusOk{
		LamportTs: 0,
	}, nil
}

func (s *Server) Publish(ctx context.Context, in *bindings.Message) (*bindings.Status, error) {
	fmt.Printf("Client %s publishing to server %s: %s", in.Sender, address, in.Contents)
	return &bindings.Status{
		LamportTs:  0,
		StatusCode: bindings.Status_OK,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := Server{}

	bindings.RegisterClientToServerServiceServer(s, &server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
