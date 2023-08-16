package main

import (
	"context"
	"fmt"
	"log"

	"github.com/RushikeshMarkad16/go-grpc-calculator/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func SumNumber(conn calculatorpb.CalculatorServiceClient) {
	fmt.Println("Initiating sum of two numbers using gRPC")

	firstNo := 3
	SecondNo := 5

	req := calculatorpb.SumRequest{
		FirstNumber: int32(firstNo),
		SecondNumer: int32(SecondNo),
	}

	resp, err := conn.Sum(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling sum function over gRPC : %v", err)
	}

	fmt.Printf("Sum of %d and %d is %d \n", firstNo, SecondNo, resp.SumResult)
}

func main() {
	//To call service methods, we first need to create a gRPC channel to communicate with the server.
	//We create this by passing the server address and port number to grpc.Dial() as follows
	cc, err := grpc.Dial(":50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect : %v", err)
	}

	defer cc.Close()

	// client stub to perform RPCs
	conn := calculatorpb.NewCalculatorServiceClient(cc)

	//Unary
	SumNumber(conn)
}
