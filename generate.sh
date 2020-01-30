#!bin/bash

# greet service
protoc greet/greetpb/greet.proto --go_out=plugins=grpc:.

# calculator service
protoc calculator/calculatorpb/calculator.proto --go_out=plugins=grpc:.