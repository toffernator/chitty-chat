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

	"github.com/toffernator/chitty-chat/client/bindings"

	"google.golang.org/grpc"
)

const (
	serverHost = "localhost:50051"
	address    = "localhost:50052"
)

type Client struct {
	serverHost string
	clientHost string
	bindings.UnimplementedClientToServerServiceServer
}

func (c *Client) Broadcast(ctx context.Context, in *bindings.Message) (*bindings.StatusOk, error) {
	fmt.Printf("Client %s received following message: %s", c.clientHost, in.Contents)
	return &bindings.StatusOk{
		LamportTs: 0,
	}, nil
}

func publish(msg string, conn *grpc.ClientConn) {
	client := bindings.NewClientToServerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client.Publish(ctx, &bindings.Message{
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

	clientServer := Client{}

	bindings.RegisterClientToServerServiceServer(s, clientServer)
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
