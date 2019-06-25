package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"grpcprac/prac1/proto/ping"
)

const (
	address     = "localhost:19003"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := ping.NewPingClient(conn)

	message := "qwer"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &ping.HelloRequest{ToMessage: message})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Ping: %s", r.ResMessage)
}
