package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/toffernator/chitty-chat/client/bindings"
	serverBindings "github.com/toffernator/chitty-chat/server/bindings"

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

func (c *Client) Join(ctx context.Context, in *bindings.Address, opts ...grpc.CallOption) (*bindings.StatusOk, error) {
	fmt.Printf("Client %s joining server %s", c.clientHost, c.serverHost)
	return &bindings.StatusOk{
		LamportTs: 0,
	}, nil
}

func (c *Client) Leave(ctx context.Context, in *bindings.Address, opts ...grpc.CallOption) (*bindings.StatusOk, error) {
	fmt.Printf("Client %s leaving server %s", c.clientHost, c.serverHost)
	return &bindings.StatusOk{
		LamportTs: 0,
	}, nil
}

func (c *Client) Publish(ctx context.Context, in *bindings.Message, opts ...grpc.CallOption) (*bindings.Status, error) {
	fmt.Printf("Client %s publishing to server %s: %s", c.clientHost, c.serverHost, in.Contents)
	return &bindings.Status{
		LamportTs:  0,
		StatusCode: bindings.Status_OK,
	}, nil
}

func main() {
	conn, err := grpc.Dial(serverHost, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	clientOfServer := serverBindings.NewServerToClientClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	message := os.Args[1]
	clientOfServer.Broadcast(ctx, &serverBindings.Message{
		LamportTs: 0,
		Contents:  message,
	})
}
