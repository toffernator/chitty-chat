package main

import (
	"context"
	"fmt"
	"time"

	"github.com/toffernator/chitty-chat/client/bindings"

	"google.golang.org/grpc"
)

const (
	serverHost = "localhost:50051"
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

func main() {
	conn, err := grpc.Dial(serverHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := bindings.NewClientToServerServiceClient(conn)
	time.Sleep(time.Duration(10) * time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	client.Publish(ctx, &bindings.Message{
		LamportTs: 0,
		Contents:  "hej",
	})
}
