package main

import (
	"context"
	"log"

	pb "example.com/go-grpc/greet/proto"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet func %v", in)
	return &pb.GreetResponse{
		Result: "hello " + in.FirstName,
	}, nil
}
