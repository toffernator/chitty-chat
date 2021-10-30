package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	chatPb "github.com/toffernator/chitty-chat/chat/protobuf"
	"github.com/toffernator/chitty-chat/logicalclock"
	notificationPb "github.com/toffernator/chitty-chat/notification/protobuf"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

type ChatServer struct {
	clients map[string]*NotificationClient
	lamport logicalclock.LamportClock
	chatPb.UnimplementedChatServiceServer
}

type NotificationClient struct {
	notificationPb.NotificationServiceClient
	Connection *grpc.ClientConn
}

func NewServer() *ChatServer {
	s := ChatServer{
		clients: make(map[string]*NotificationClient),
	}
	return &s
}

func (s *ChatServer) Join(ctx context.Context, in *chatPb.Address) (*chatPb.StatusOk, error) {
	fmt.Printf("Client %s joining server %s", in.Address, address)

	// Create a new client for the given address
	conn, err := grpc.Dial(in.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Client with address %s failed to join\n", in.Address)
	}

	// Adding to map
	client := NotificationClient{
		notificationPb.NewNotificationServiceClient(conn),
		conn,
	}
	s.clients[in.Address] = &client

	broadcastMsg := fmt.Sprintf("%s @ %d joined", in.Address, s.lamport.Read())
	s.broadcast(broadcastMsg)

	return &chatPb.StatusOk{
		LamportTs: 0,
	}, nil
}

func (s *ChatServer) Leave(ctx context.Context, in *chatPb.Address) (*chatPb.StatusOk, error) {
	log.Printf("Client %s leaving", in.Address)

	for address, client := range s.clients {
		if address == in.Address {
			client.Connection.Close()
			delete(s.clients, address)
		}
	}

	broadcastMsg := fmt.Sprintf("%s has left @ %d", in.Address, s.lamport.Read())
	s.broadcast(broadcastMsg)

	return &chatPb.StatusOk{}, nil
}

func (s *ChatServer) Publish(ctx context.Context, in *chatPb.Message) (*chatPb.Status, error) {
	fmt.Printf("Client %s publishing to server %s: %s", in.Sender, address, in.Contents)
	broadcastMsg := fmt.Sprintf("%s @ %d: %s", in.Sender, s.lamport.Read(), in.Contents)
	s.broadcast(broadcastMsg)
	return &chatPb.Status{
		LamportTs:  0,
		StatusCode: chatPb.Status_OK,
	}, nil
}

func (s *ChatServer) broadcast(msg string) {
	for _, client := range s.clients {
		log.Printf("%s\n", client.Connection.Target())

		requestctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		_, err := client.Notify(requestctx, &notificationPb.Message{
			LamportTs: 0,
			Contents:  msg,
		})
		if err != nil {
			log.Println(err.Error())
		}
	}
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := NewServer()

	chatPb.RegisterChatServiceServer(s, server)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
