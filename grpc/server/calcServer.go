package main

import (
	"context"

	pb "github.com/brianreynaldgit/golang_playground"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) SumMethod1(ctx context.Context)
