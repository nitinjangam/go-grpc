package main

import (
	"context"
	"log"

	"github.com/nitinjangam/go-grpc/QuoteService/Quotepb"
	"google.golang.org/grpc"
)

func main() {
	log.Println("Hello I am Client")
	cc, err := grpc.Dial("Localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := Quotepb.NewQuoteServiceClient(cc)

	doUnary(c)
}

func doUnary(c Quotepb.QuoteServiceClient) {
	req := Quotepb.QuoteRequest{
		GetQoute: "Hi",
	}
	res, err := c.Quote(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error while calling Quote RPC: %v", err)
	}
	log.Printf("Response from Qoute RPC: %v", res)

}
