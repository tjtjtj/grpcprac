package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"golang.org/x/net/context"
	"grpcprac/prac1/proto/ping"
)

type PingService struct {
}

func (s *PingService) Hello(ctx context.Context, req *ping.HelloRequest) (*ping.HelloResponse, error) {
	toMessage := req.GetToMessage()
	log.Println(toMessage)
	response := ping.HelloResponse{
		ResMessage: "I hear " + toMessage,
	}
	return &response, nil
}

var port = ":19003"

func main() {
	listenPort, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalln(err)
	}
	server := grpc.NewServer()
	pingService := PingService{}
	ping.RegisterPingServer(server, &pingService)

	log.Printf("[server started] localhost%s", port)
	err = server.Serve(listenPort)
	if err != nil {
		log.Fatalln(err)
	}
}
