package main

import (
	"context"
	"log"
	"net"

	pb "../protobuf"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) TesteMaroto(ctx context.Context, in *pb.ObjetoEntrada) (*pb.ObjetoSaida, error) {
	log.Printf("Received: %v", in.Nome)
	return &pb.ObjetoSaida{Nome: in.Nome}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMarotoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
