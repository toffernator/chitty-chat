package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/toffernator/chitty-chat/server/bindings"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type Server struct {
	clients []string
	bindings.UnimplementedServerToClientServer
}

func (s *Server) Broadcast(ctx context.Context, in *bindings.Message) (*bindings.StatusOk, error) {
	for _, client := range s.clients {
		fmt.Printf("Broadcasting to client %s: %s", client, in.Contents)
	}
	return &bindings.StatusOk{}, nil
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := Server{}

	bindings.RegisterServerToClientServer(s, &server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
