package main

import (
	"context"
	"fmt"
	"log"

	"github.com/victor-nach/grpc-intro/calculator/calculatorpb"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Calculator client...")

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}

	fmt.Println("successfuly connected..")
	defer conn.Close()

	// create a client stub
	c := calculatorpb.NewCalculatorServiceClient(conn)

	doAdd(c)
}

func doAdd(c) {
	req := &calculatorpb.AddRequest{
		FirstNum:  5,
		SecondNum: 6,
	}
	res, err := c.Add(context.Background(), req)
	if err != nil {
		log.Fatalf("Error occured while calling add: %v", err)
	}
	log.Println("Response from add service: %v", res.Sum)
}
