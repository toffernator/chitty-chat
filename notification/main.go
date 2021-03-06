package main

import (
	"bufio"
	"context"
	"log"
	"net"
	"os"
	"strings"
	"time"

	chatPb "github.com/toffernator/chitty-chat/chat/protobuf"
	"github.com/toffernator/chitty-chat/logicalclock"
	notificationPb "github.com/toffernator/chitty-chat/notification/protobuf"

	"google.golang.org/grpc"
)

type NotificationServer struct {
	address string
	lamport logicalclock.LamportTimer
	notificationPb.UnimplementedNotificationServiceServer
}

var (
	address string
	lamport logicalclock.LamportTimer
)

type ChatClient struct {
	chatPb.ChatServiceClient
	lamport logicalclock.LamportTimer
	conn    *grpc.ClientConn
}

func (n *NotificationServer) Broadcast(ctx context.Context, in *notificationPb.Message) (*notificationPb.StatusOk, error) {
	log.Printf("Received a message at Lamport time %d: %s\n", n.lamport.Read(), in.Contents)
	n.lamport.Update(logicalclock.NewLamportClock(in.LamportTs))

	return &notificationPb.StatusOk{
		LamportTs: n.lamport.Read(),
	}, nil
}

func join(target string) *ChatClient {
	conn, err := grpc.Dial(target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}

	client := &ChatClient{
		chatPb.NewChatServiceClient(conn),
		lamport,
		conn,
	}

	client.lamport.Increment()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	status, err := client.Join(ctx, &chatPb.Address{
		LamportTs: client.lamport.Read(),
		Address:   address,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	client.lamport.Update(logicalclock.NewLamportClock(status.LamportTs))

	return client
}

func (c *ChatClient) publish(msg string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c.lamport.Increment()

	status, err := c.Publish(ctx, &chatPb.Message{
		LamportTs: c.lamport.Read(),
		Sender:    address,
		Contents:  msg,
	})
	if err != nil {
		log.Fatalf(err.Error())
	}
	if status.StatusCode == 1 {
		log.Println("ChatServer rejected the message since it was invalid.")
	} else if status.StatusCode == 2 {
		log.Println("ChatServer was not able to broadcast to all clients.")
	}
	c.lamport.Update(logicalclock.NewLamportClock(status.LamportTs))
}

func (c *ChatClient) leave() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	defer c.conn.Close()

	c.lamport.Increment()

	status, err := c.Leave(ctx, &chatPb.Address{Address: address})
	if err != nil {
		log.Fatalf(err.Error())
	}

	c.lamport.Update(logicalclock.NewLamportClock(status.LamportTs))
	log.Println("Disconnected.")
}

func serve() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()

	notificationServer := &NotificationServer{
		address: address,
		lamport: lamport,
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
	if len(os.Args) < 2 {
		log.Fatalln("Must pass the listen address as an argument")
	}
	address = os.Args[1]
	lamport = logicalclock.NewLamportClock(0)

	go handleUserInput()
	serve()
}
