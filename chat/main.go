package main

import (
	"context"
	"errors"
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
	lamport logicalclock.LamportTimer
	chatPb.UnimplementedChatServiceServer
}

type NotificationClient struct {
	notificationPb.NotificationServiceClient
	Connection *grpc.ClientConn
}

func NewServer() *ChatServer {
	s := ChatServer{
		clients: make(map[string]*NotificationClient),
		lamport: logicalclock.NewLamportClock(0),
	}
	return &s
}

func (s *ChatServer) Join(ctx context.Context, in *chatPb.Address) (*chatPb.StatusOk, error) {
	log.Printf("Participant %s joined Chitty-Chat at Lamport time %d\n", in.Address, s.lamport.Read())
	s.lamport.Update(logicalclock.NewLamportClock(in.LamportTs))

	conn, err := grpc.Dial(in.Address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Printf("Participant %s failed to join Chitty-Chat at Lamport time %d\n", in.Address, s.lamport.Read())
		return nil, errors.New("Server failed to establish a notification connection")
	}

	client := NotificationClient{
		notificationPb.NewNotificationServiceClient(conn),
		conn,
	}
	s.clients[in.Address] = &client

	broadcastMsg := fmt.Sprintf("Participant %s joined Chitty-Chat at Lamport time %d\n", in.Address, s.lamport.Read())
	s.broadcast(broadcastMsg)

	s.lamport.Increment()

	return &chatPb.StatusOk{
		LamportTs: s.lamport.Read(),
	}, nil
}

func (s *ChatServer) Leave(ctx context.Context, in *chatPb.Address) (*chatPb.StatusOk, error) {
	log.Printf("Participant %s left Chitty-Chat at Lamport time %d\n", in.Address, s.lamport.Read())
	s.lamport.Update(logicalclock.NewLamportClock(in.LamportTs))

	for address, client := range s.clients {
		if address == in.Address {
			if err := client.Connection.Close(); err != nil {
				log.Fatalf("Failed to close connection with error: %s\n", err.Error())
			}

			delete(s.clients, address)
		}
	}

	broadcastMsg := fmt.Sprintf("Participant %s left Chitty-Chat at Lamport time %d\n", in.Address, s.lamport.Read())
	s.broadcast(broadcastMsg)

	s.lamport.Increment()
	return &chatPb.StatusOk{
		LamportTs: s.lamport.Read(),
	}, nil
}

func (s *ChatServer) Publish(ctx context.Context, in *chatPb.Message) (*chatPb.Status, error) {
	s.lamport.Update(logicalclock.NewLamportClock(in.LamportTs))
	if validateMessage(in.Contents) {
		log.Printf("Ts: %d -- Client %s publishing to server %s: %s\n", s.lamport.Read(), in.Sender, address, in.Contents)
	} else {
		log.Printf("Ts: %d -- Client %s attempted publishing an invalid message.", s.lamport.Read(), in.Sender)
		s.lamport.Increment()
		return &chatPb.Status{
			LamportTs:  s.lamport.Read(),
			StatusCode: chatPb.Status_INVALIDMSG,
		}, nil
	}

	broadcastMsg := fmt.Sprintf("%s: %s", in.Sender, in.Contents)
	nodes, successes := s.broadcast(broadcastMsg)
	s.lamport.Increment()
	if successes < nodes {
		log.Printf("Ts: %d -- Attempted to broadcast to %d nodes but only %d requests succeeded", s.lamport.Read(), nodes, successes)
		return &chatPb.Status{
			LamportTs:  s.lamport.Read(),
			StatusCode: chatPb.Status_INCOMPLETEBROADCAST,
		}, nil
	} else {
		return &chatPb.Status{
			LamportTs:  s.lamport.Read(),
			StatusCode: chatPb.Status_OK,
		}, nil
	}

}

func (s *ChatServer) broadcast(msg string) (total int, successes int) {
	failed := 0
	for address, client := range s.clients {
		s.lamport.Increment()
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()

		_, err := client.Broadcast(ctx, &notificationPb.Message{
			LamportTs: s.lamport.Read(),
			Contents:  "ChatServer " + msg,
		})
		if err != nil {
			log.Printf("Ts: %d -- Failed to Broadcast client at %s with error: %v\n", s.lamport.Read(), address, err.Error())
			failed++
		}
	}
	return len(s.clients), len(s.clients) - failed
}

func validateMessage(msg string) bool {
	return len(msg) <= 128
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	server := NewServer()

	chatPb.RegisterChatServiceServer(s, server)
	log.Printf("server listening at %v\n", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
