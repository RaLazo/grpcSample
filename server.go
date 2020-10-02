package main

import (
	"log"
	"net"

	"github.com/ralazo/grpcSample/chat"
	"google.golang.org/grpc"
)

func main() {
	// create a TCP Listner on Port 9000
	lis, err := net.Listen("tcp", ":9000")
	// this how you handle errors in Golang
	if err != nil {
		log.Fatalf("Failed to listen on port 9000 %v", err)
	}

	// this is just a structure that has an interface with needed function SayHello
	// have a look at /chat/chat.go for more information
	s := chat.Server{}

	//create the GRCP Server
	grpcServer := grpc.NewServer()

	//
	chat.RegisterChatServiceServer(grpcServer, &s)

	// start listening on port 9000 for rpc
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server over port 9000 %v", err)
	}

}
