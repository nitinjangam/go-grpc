package main

import (
	"context"
	"log"
	"time"

	messageService "github.com/nitinjangam/go-grpc/MessageService/messageService"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello I am Client")
	cc, err := grpc.Dial("Localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := messageService.NewMessageServiceClient(cc)

	doUnary(c)
}

func doUnary(c messageService.MessageServiceClient) {
	req := messageService.SendMessageRequest{
		Greet: "Hi",
	}
	t1 := time.Now()
	res, err := c.SendMessage(context.Background(), &req)
	t2 := time.Now()
	diff := t2.Sub(t1)
	log.Printf("Time taken: %v", diff)
	if err != nil {
		log.Fatalf("error while calling SendMessage RPC: %v", err)
	}
	log.Printf("Response from SendMessage: %v", res.Message)
}
