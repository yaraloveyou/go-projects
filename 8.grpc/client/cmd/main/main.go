package main

import (
	"context"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc/server/proto"
	"log"
	"time"
)

var (
	addr = flag.String("addr", "localhost:50051", "server address")
	name = flag.String("name", "World", "Value")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewHelloServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Value: *name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Gretting: %s", r.GetValue())
}
