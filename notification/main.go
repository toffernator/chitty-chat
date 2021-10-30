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

type ChatClient struct {
	chatPb.ChatServiceClient
	conn *grpc.ClientConn
}

func (n *NotificationServer) Notify(ctx context.Context, in *notificationPb.Message) (*notificationPb.StatusOk, error) {
	fmt.Printf("Client %s received following message: %s", n.clientHost, in.Contents)
	return &notificationPb.StatusOk{
		LamportTs: 0,
	}, nil
}

func join(address string) *ChatClient {
	conn, err := grpc.Dial(serverHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := &ChatClient{
		chatPb.NewChatServiceClient(conn),
		conn,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.Join(ctx, &chatPb.Address{Address: "localhost:50052"})
	if err != nil {
		log.Fatalf(err.Error())
	}

	return client
}

func (c *ChatClient) publish(msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err := c.Publish(ctx, &chatPb.Message{
		LamportTs: 0,
		Sender:    address,
		Contents:  msg,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func (c *ChatClient) leave() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	defer c.conn.Close()

	_, err := c.Leave(ctx, &chatPb.Address{Address: address})
	if err != nil {
		log.Fatalf(err.Error())
	}
}

func serve() {
	// Set-up client "service"
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	notificationServer := &NotificationServer{
		serverHost: serverHost,
		clientHost: address,
	}

	notificationPb.RegisterNotificationServiceServer(s, notificationServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handleUserInput() {
	client := join(serverHost)
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			parsed := strings.SplitN(input, " ", 2)

			switch parsed[0] {
			case "/publish":
				client.publish(parsed[1])
			case "/leave":
				client.leave()
				os.Exit(0)
			}
		}
	}
}

func main() {
	go handleUserInput()
	serve()
}
