package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/victor-nach/grpc-intro/greet/greetpb"
)

type server struct{}

func main() {
	fmt.Println("hello world!")

	// write a listener, then assign it to a variable
	// 50051 is the default port for grpc
	lis, err := net.Listen("tcp", "0.0.0.0:50051") // this is the port binding
	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	// create a new grpc server
	s := grpc.NewServer()
	greetpb.RegisterGreetServiceServer(s, &server{}) // here we can register as many services as we want

	// bind the port to a grpc server
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}