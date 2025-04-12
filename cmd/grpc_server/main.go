package main

import (
	"context"
	"fmt"
	"log"
	"net"

	desc "github.com/armanbektassov/go_chat/pkg/chat_v1"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	desc.UnimplementedChatV1Server
}

// Create ...
func (s *server) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	log.Printf("Chat id: %s", req.GetUsernames())

	return &desc.CreateResponse{
		Id: gofakeit.Int64(),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
