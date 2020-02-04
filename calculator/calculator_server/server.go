package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/victor-nach/grpc-intro/calculator/calculatorpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Add(ctx context.Context, req *calculatorpb.AddRequest) (*calculatorpb.AddResponse, error) {
	firstNum := req.GetFirstNum()
	secondNum := req.GetSecondNum()
	sum := firstNum + secondNum
	res := &calculatorpb.AddResponse{
		Sum: sum,
	}

	return res, nil
}
func main() {
	fmt.Println("calculator server...")
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}

	fmt.Println("Serving calculator service...")
}
