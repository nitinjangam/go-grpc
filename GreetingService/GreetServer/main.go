package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/nitinjangam/go-grpc/GreetingService/Greetpb"
	"google.golang.org/grpc"
)

type server struct {
	Greetpb.UnimplementedGreetingServiceServer
}

func (*server) Greet(ctx context.Context, req *Greetpb.Greeting) (*Greetpb.Greeting, error) {
	log.Printf("Greet Function was invoked with %v", req)
	currentTime := time.Now()
	timeStampString := currentTime.Format("2006-01-02 15:04:05")
	layOut := "2006-01-02 15:04:05"
	timeStamp, err := time.Parse(layOut, timeStampString)
	if err != nil {
		fmt.Println(err)
	}
	hr, _, _ := timeStamp.Clock()
	wish := ""
	fmt.Println(hr)
	if 5 < hr && hr < 12 {
		wish = "Good Morning"
	} else if 12 < hr && hr < 17 {
		wish = "Good After Noon"
	} else if hr >= 17 {
		wish = "Good Evening"
	}
	res := Greetpb.Greeting{
		GreetMessage: wish,
	}
	return &res, nil

}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	Greetpb.RegisterGreetingServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
