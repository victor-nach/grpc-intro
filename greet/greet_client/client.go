package main

import (
	"context"
	"fmt"
	"log"

	"github.com/victor-nach/grpc-intro/greet/greetpb"
	"google.golang.org/grpc"
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

	// create a client for the greet service
	c := greetpb.NewGreetServiceClient(conn)
	// fmt.Printf("Created client: %f", c)
	doUnary(c)

}

func doUnary(c greetpb.GreetServiceClient) {
	fmt.Println("Starting to do Unary RPC....")
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Iheanacho",
			LastName:  "Victor",
		},
	}

	// the created client has access to the calls defined in the protobuf file
	// so here we do a rpc call to the function as defined in the service by passing in a greetRequest type
	// then we get a GreetResponse type as defined in the protobuf file
	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling greet RPC: %v\n", err)
	}
	log.Printf("Response from Greet: %v", res.Result)
}