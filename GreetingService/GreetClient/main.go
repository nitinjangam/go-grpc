package main

import (
	"context"
	"log"

	"github.com/nitinjangam/go-grpc/GreetingService/Greetpb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello I am Client")
	cc, err := grpc.Dial("Localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := Greetpb.NewGreetingServiceClient(cc)

	doUnary(c)
}

func doUnary(c Greetpb.GreetingServiceClient) {
	req := Greetpb.Greeting{
		GreetMessage: "Hello",
	}

	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet RPC: %v", res)
}
