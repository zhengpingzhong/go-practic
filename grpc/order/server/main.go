package main

import (
	"google.golang.org/grpc"
	"log"
	"net"
	pb "zhengpingzhong.com/go-practic/grpc/oder/server/ecommerce"
)

const (
	port = ":8888"
)

type server struct {
	pb.UnimplementedOrderManagementServer
}

var orderMap map[string]*pb.Order

//func (s *server) GetOrder(ctx context.Context, orderId *wrapperspb.StringValue) (*pb.Order, error) {
//	ord := orderMap[orderId.Value]
//	fmt.Println(ord)
//
//}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: #{err}")
	}

	newServer := grpc.NewServer()

	pb.RegisterOrderManagementServer(newServer, &server{})

	if err := newServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: #{err}")
	}

}
