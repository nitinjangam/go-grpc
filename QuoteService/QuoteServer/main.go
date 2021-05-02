package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	"github.com/nitinjangam/go-grpc/QuoteService/Quotepb"
	"google.golang.org/grpc"
)

type server struct {
	Quotepb.UnimplementedQuoteServiceServer
}

var quotes = []string{
	"You Learn More From Failure Than From Success. Don’t Let It Stop You.",
	"Don’t Let Yesterday Take Up Too Much Of Today.",
	"Failure Will Never Overtake Me If My Determination To Succeed Is Strong Enough.",
	"We May Encounter Many Defeats But We Must Not Be Defeated.",
	"Knowing Is Not Enough; We Must Apply. Wishing Is Not Enough; We Must Do.",
	"The Man Who Has Confidence In Himself Gains The Confidence Of Others.",
	"Creativity Is Intelligence Having Fun.",
}

func (*server) Quote(ctx context.Context, req *Quotepb.QuoteRequest) (*Quotepb.QuoteResponse, error) {
	log.Printf("Quote function was invoked with %v", req)
	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	fmt.Println(y1.Intn(len(quotes)))
	quote := quotes[y1.Intn(len(quotes))]
	res := Quotepb.QuoteResponse{
		QuoteReply: "Quote of the day: " + quote,
	}
	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	Quotepb.RegisterQuoteServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
