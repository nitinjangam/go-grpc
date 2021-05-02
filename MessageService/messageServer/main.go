package main

import (
	"context"
	"log"
	"net"
	"sync"

	"github.com/nitinjangam/go-grpc/GreetingService/Greetpb"
	messageService "github.com/nitinjangam/go-grpc/MessageService/messageService"
	"github.com/nitinjangam/go-grpc/QuoteService/Quotepb"
	"google.golang.org/grpc"
)

type server struct {
	messageService.UnimplementedMessageServiceServer
}

func (*server) SendMessage(ctx context.Context, req *messageService.SendMessageRequest) (*messageService.SendMessageResponse, error) {
	wg := &sync.WaitGroup{}
	log.Printf("SendMessage Function was invoked with %v", req)
	greeting := req.GetGreet()

	// Calling Greet Service
	var resp *Greetpb.Greeting
	var resp1 *Quotepb.QuoteResponse
	greet := func() {
		defer wg.Done()
		cc, err := grpc.Dial("Localhost:50052", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}
		c := Greetpb.NewGreetingServiceClient(cc)

		resp, err = callGreetService(c)
		if err != nil {
			log.Fatalf("Error while calling GreetService: %v", err)
		}
		cc.Close()
	}
	wg.Add(1)
	go greet()

	// Calling QuoteService
	quote := func() {
		defer wg.Done()
		cc1, err := grpc.Dial("Localhost:50053", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("Could not connect: %v", err)
		}

		c1 := Quotepb.NewQuoteServiceClient(cc1)

		resp1, err = callQuoteService(c1)
		if err != nil {
			log.Fatalf("Error while calling QuoteService;%v", err)
		}
		cc1.Close()
	}
	wg.Add(1)
	go quote()

	wg.Wait()

	result := &messageService.SendMessageResponse{
		Message: greeting + ", " + resp.GetGreetMessage() + ".\n" + resp1.GetQuoteReply(),
	}
	return result, nil
}

func callGreetService(c Greetpb.GreetingServiceClient) (*Greetpb.Greeting, error) {
	req := Greetpb.Greeting{
		GreetMessage: "Hello",
	}

	res, err := c.Greet(context.Background(), &req)
	if err != nil {
		log.Fatalf("error while calling Greet RPC: %v", err)
	}
	log.Printf("Response from Greet RPC: %v", res)
	return res, nil
}

func callQuoteService(c Quotepb.QuoteServiceClient) (*Quotepb.QuoteResponse, error) {
	req := Quotepb.QuoteRequest{
		GetQoute: "Hi",
	}
	res, err := c.Quote(context.Background(), &req)
	if err != nil {
		log.Fatalf("Error while calling Quote RPC: %v", err)
	}
	log.Printf("Response from Qoute RPC: %v", res)
	return res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	messageService.RegisterMessageServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
