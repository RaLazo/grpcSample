package main

import (
	"log"

	"golang.org/x/net/context"

	"github.com/ralazo/grpcSample/chat"
	"google.golang.org/grpc"
)

func main() {
	// create client for GRPC Server
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("no server connection could be established cause: %v", err)
	}

	// defer runs after the functions finishes
	defer conn.Close()

	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "Hello from the client!",
	}

	response, err := c.SayHello(context.Background(), &message)
	log.Printf("Response from Server: %s", response.Body)
}
