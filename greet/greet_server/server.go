package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/victor-nach/grpc-intro/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("Greet function was invoked with:  %v\n", req)
	firstName := req.GetGreeting().GetFirstName()
	result := "Hello " + firstName
	res := &greetpb.GreetResponse{
		Result: result,
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) (error) {
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 10; i++ {
		result := "Hello " + firstName + ": " + strconv.Itoa(i)
		res := &greetpb.GreetManyTimesResponse{
			Result: result,
		}
		stream.Send(res)
		time.Sleep(1000 * time.Millisecond)
	}
	return nil
}

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

	// here we can register as many services as we want on the server we created
	// already from the proto file we defined, the greetService has a rpc call 'Greet'
	// so the struct that we pass to this must implement that method
	greetpb.RegisterGreetServiceServer(s, &server{})

	// we already have a listener
	// so we serve the server that already has all our services attached using that listener
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
