package main

import (
	"context"
	"log"
	"time"

	pb "github.com/brianreynaldgit/golang_playground/grpc/calculator"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	client := pb.NewCalculatorClient(conn)
	sumMethod1(client, 30, 20)
	var s []int32 = []int32{2, 3, 5, 7, 11, 13}
	sumMethod2(client, s)
}

func sumMethod1(client pb.CalculatorClient, A int32, B int32) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SumMethod1(ctx, &pb.CalcInputs1{
		A: A,
		B: B,
	})
	if err != nil {
		log.Fatalf("could not hit: %v", err)
	}
	log.Printf("sum method 1: %v", r.GetOut())
}
func sumMethod2(client pb.CalculatorClient, listInput []int32) {
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := client.SumMethod2(ctx, &pb.CalcInputs2{
		Inputs: listInput,
	})
	if err != nil {
		log.Fatalf("could not hit: %v", err)
	}
	log.Printf("sum method 1: %v", r.GetOut())
}
