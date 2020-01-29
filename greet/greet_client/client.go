package main

import (
	"fmt"
	"log"
	"google.golang.org/grpc"
	"github.com/victor-nach/grpc-intro/greet/greetpb"

)

func main() {
	fmt.Println("hello, I'm a client")

	// create a connection to the server
	// tahes two option, the address and some options
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	defer conn.Close()

	// create a client
	c := greetpb.NewGreetServiceClient(conn)
	fmt.Printf("Created client: %f", c)
}