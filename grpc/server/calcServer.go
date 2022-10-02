package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/brianreynaldgit/golang_playground/grpc/calculator"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

var (
	port = flag.Int("port", 50051, "The server port")
)

func (s *server) SumMethod1(ctx context.Context, in *pb.CalcInputs1) (*pb.CalcOutput, error) {
	sum := in.A + in.B
	out := &pb.CalcOutput{Out: sum}

	return out, nil
}

func (s *server) SumMethod2(ctx context.Context, in *pb.CalcInputs2) (*pb.CalcOutput, error) {
	var sum int32
	for _, data := range in.Inputs {
		sum += data
	}
	out := &pb.CalcOutput{Out: sum}

	return out, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
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
