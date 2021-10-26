package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	ClientBindings "github.com/toffernator/chitty-chat/client/bindings"
	ServerBindings "github.com/toffernator/chitty-chat/server/bindings"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type Server struct {
	clients []ServerBindings.ServerToClientClient
	ClientBindings.UnimplementedClientToServerServiceServer
}

func (s *Server) Join(ctx context.Context, in *ClientBindings.Address) (*ClientBindings.StatusOk, error) {
	fmt.Printf("Client %s joining server %s", in.Address, address)
	return &ClientBindings.StatusOk{
		LamportTs: 0,
	}, nil
}

func (s *Server) Leave(ctx context.Context, in *ClientBindings.Address) (*ClientBindings.StatusOk, error) {
	fmt.Printf("Client %s leaving server %s", in.Address, address)
	return &ClientBindings.StatusOk{
		LamportTs: 0,
	}, nil
}

func (s *Server) Publish(ctx context.Context, in *ClientBindings.Message) (*ClientBindings.Status, error) {
	fmt.Printf("Client %s publishing to server %s: %s", in.Sender, address, in.Contents)
	s.broadcast(in.Contents)
	return &ClientBindings.Status{
		LamportTs:  0,
		StatusCode: ClientBindings.Status_OK,
	}, nil
}

func (s *Server) broadcast(msg string) {
	for _, client := range s.clients {
		requestctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		go client.Broadcast(requestctx, &ServerBindings.Message{
			LamportTs: 0,
			Contents:  msg,
		})
	}
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := Server{}

	ClientBindings.RegisterClientToServerServiceServer(s, &server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
