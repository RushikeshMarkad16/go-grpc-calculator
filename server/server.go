package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/RushikeshMarkad16/go-grpc-calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	calculatorpb.CalculatorServiceServer
}

func (*server) Sum(ctx context.Context, req *calculatorpb.SumRequest) (resp *calculatorpb.SumResponse, err error) {
	fmt.Println("The Sum function is called via gRPC")

	//get Req param
	firstNumber := req.FirstNumber
	secondNumber := req.SecondNumer

	//execute the business logic
	result := firstNumber + secondNumber

	//prepare the response
	resp = &calculatorpb.SumResponse{
		SumResult: result,
	}

	return resp, nil
}

func main() {
	fmt.Println("Starting grpc server....")

	//Specify the port we want to use to listen for client requests
	listen, err := net.Listen("tcp", ":50051")

	if err != nil {
		log.Fatalf("Failed to listen : %v", err)
	}

	//Create an instance of the gRPC server
	s := grpc.NewServer()

	//Register our service implementation with the gRPC server.
	calculatorpb.RegisterCalculatorServiceServer(s, &server{})

	//Register reflection service on gRPC server
	reflection.Register(s)

	if err := s.Serve(listen); err != nil {
		log.Fatalf("Failed to serve : %v", err)
	}
}
