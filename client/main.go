package main

import (
	"context"
	"fmt"

	"github.com/toffernator/chitty-chat/client/bindings"

	"google.golang.org/grpc"
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

}
