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

type NotificationServer struct {
	address string
	notificationPb.UnimplementedNotificationServiceServer
}

var (
	address string
)

type ChatClient struct {
	chatPb.ChatServiceClient
	conn *grpc.ClientConn
}

func (n *NotificationServer) Notify(ctx context.Context, in *notificationPb.Message) (*notificationPb.StatusOk, error) {
	fmt.Printf("Client %s received following message: %s\n", n.address, in.Contents)
	return &notificationPb.StatusOk{
		LamportTs: 0,
	}, nil
}

func join(target string) *ChatClient {
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := &ChatClient{
		chatPb.NewChatServiceClient(conn),
		conn,
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err = client.Join(ctx, &chatPb.Address{Address: address})
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
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	notificationServer := &NotificationServer{
		address: address,
	}

	notificationPb.RegisterNotificationServiceServer(s, notificationServer)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func handleUserInput() {
	var client *ChatClient
	for {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			input := scanner.Text()
			parsed := strings.SplitN(input, " ", 2)

			switch parsed[0] {
			case "/join":
				if client == nil {
					client = join(parsed[1])
				} else {
					log.Println("You must call /leave before joining a new chat service")
				}
			case "/publish":
				if client != nil {
					client.publish(parsed[1])
				} else {
					log.Println("You must call /join before publishing")
				}
			case "/leave":
				if client != nil {
					client.leave()
					client = nil
				} else {
					log.Println("You must call /join before leaving")
				}
			}
		}
	}
}

func main() {
	address = os.Args[1]

	go handleUserInput()
	serve()
}
