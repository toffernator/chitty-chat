package main

import (
	"fmt"

	"github.com/toffernator/chitty-chat/server/bindings"
)

type Server struct {
	clients []string
	bindings.UnimplementedServerToClientServer
}

func (s *Server) Broadcast(msg string) {
	for _, client := range s.clients {
		fmt.Printf("Broadcasting to client %s: %s", client, msg)
	}
}

func main() {

}
