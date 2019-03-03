package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "../protobuf"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "piroca"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewMarotoClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	r, err := c.TesteMaroto(ctx, &pb.ObjetoEntrada{
		Nome:  name,
		Idade: "1",
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.Nome)
}
