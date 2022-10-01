package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/brianreynaldgit/golang_playground/grpc/calculator"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) SumMethod1(ctx context.Context, in *pb.CalcInputs1) (*pb.CalcOutput, error) {
	var output *pb.CalcOutput

	output.Out = in.A + in.B

	return output, nil
}

func (s *server) SumMethod2(ctx context.Context, in *pb.CalcInputs2) (*pb.CalcOutput, error) {
	var output *pb.CalcOutput

	for _, data := range in.Inputs {
		output.Out += data
	}

	return output, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
