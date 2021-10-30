package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"

	chatPb "github.com/toffernator/chitty-chat/chat/protobuf"
	notificationPb "github.com/toffernator/chitty-chat/notification/protobuf"

	"google.golang.org/grpc"
)

const (
	serverHost = "localhost:50051"
	address    = "localhost:50052"
)

type NotificationServer struct {
	serverHost string
	clientHost string
	notificationPb.UnimplementedNotificationServiceServer
}

func (n *NotificationServer) Notify(ctx context.Context, in *notificationPb.Message) (*notificationPb.StatusOk, error) {
	fmt.Printf("Client %s received following message: %s", n.clientHost, in.Contents)
	return &notificationPb.StatusOk{
		LamportTs: 0,
	}, nil
}

func publish(msg string, conn *grpc.ClientConn) {
	client := chatPb.NewChatServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client.Publish(ctx, &chatPb.Message{
		LamportTs: 0,
		Sender:    address,
		Contents:  msg,
	})
}

func serve() {
	// Set-up client "service"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	notificationServer := NotificationServer{}

	notificationPb.RegisterNotificationServiceServer(s, &notificationServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	go serve()

	// Connect to remote server
	conn, err := grpc.Dial(serverHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// Parse user-input
	for {
		// Read user input
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()

			// Parse input
			parsed := strings.SplitN(input, " ", 2)
			fmt.Printf("%v\n", parsed)

			// Act on user input
			switch parsed[0] {
			case "/publish":
				publish(parsed[1], conn)
			}
		}
	}
}
