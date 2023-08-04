package main

import (
	"log"
	"net"

	pb "example.com/go-grpc/greet/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("Fail to listen %v", err)
	}

	log.Printf("listening on %v", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Fail to Serve %v", err)
	}
}
