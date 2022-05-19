package main

import (
	"context"
	pb "demo.com/productinfo/ecommerce"
	"flag"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	flag.Parse()
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewProductInfoClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.AddProduct(ctx, &pb.Product{Name: "testProductName"})
	if err != nil {
		log.Fatalf("could not add product: %v", err)
	}
	log.Printf("add product:%v", r.GetValue())
}
