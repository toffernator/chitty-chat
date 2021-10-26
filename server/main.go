package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	ClientBindings "github.com/toffernator/chitty-chat/client/bindings"
	"github.com/toffernator/chitty-chat/logicalclock"
	ServerBindings "github.com/toffernator/chitty-chat/server/bindings"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type Server struct {
	clients map[string]ServerBindings.ServerToClientClient
	lamport logicalclock.LamportClock
	ClientBindings.UnimplementedClientToServerServiceServer
}

func NewServer() *Server {
	s := Server{
		clients: make(map[string]ServerBindings.ServerToClientClient),
	}
	return &s
}

func (s *Server) Join(ctx context.Context, in *ClientBindings.Address) (*ClientBindings.StatusOk, error) {
	fmt.Printf("Client %s joining server %s", in.Address, address)

	// Create a new client for the given address
	conn, err := grpc.Dial(in.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Client with address %s failed to join\n", in.Address)
	}
	// USUALLY HERE: defer conn.close()

	client := ServerBindings.NewServerToClientClient(conn)
	s.clients = append(s.clients, client)
	broadcastMsg := fmt.Sprintf("%s @ % joined", in.Address, s.lamport.Read())
	s.broadcast(broadcastMsg)

	return &ClientBindings.StatusOk{
		LamportTs: 0,
	}, nil
}

func (s *Server) Leave(ctx context.Context, in *ClientBindings.Address) (*ClientBindings.StatusOk, error) {
	log.Printf("Client %s leaving", in.Address)

	for _, client := range s.clients {

	}
}

func (s *Server) Publish(ctx context.Context, in *ClientBindings.Message) (*ClientBindings.Status, error) {
	fmt.Printf("Client %s publishing to server %s: %s", in.Sender, address, in.Contents)
	broadcastMsg := fmt.Sprintf("%s @ %d: %s", in.Sender, s.lamport.Read(), in.Contents)
	s.broadcast(broadcastMsg)
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

	server := NewServer()

	ClientBindings.RegisterClientToServerServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
