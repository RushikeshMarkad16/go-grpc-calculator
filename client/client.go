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
	cc, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Cannot connect : %v", err)
	}

	defer cc.Close()

	conn := calculatorpb.NewCalculatorServiceClient(cc)

	//Unary
	SumNumber(conn)
}
