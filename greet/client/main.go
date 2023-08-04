package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "example.com/go-grpc/greet/proto"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("client Failed %v", err)
	}
	defer conn.Close()
	log.Fatalf("this is running: %s", "quang is running good ")

	c := pb.NewGreetServiceClient(conn)
	log.Fatalf("c")

	greetReponse, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "quang",
	})
	if err != nil {
		log.Fatalf("Greet error: %s", err)

	}
	log.Printf("greetReponse: %s", greetReponse.Result)

}
