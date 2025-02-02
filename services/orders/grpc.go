package main

import (
	"log"
	"net"

	handler "github.com/DivyanshuShekhar55/go-grpc/services/orders/handler/orders"
	"github.com/DivyanshuShekhar55/go-grpc/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {

	// listen to a tcp server, via listener
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen %v", err)
	}

	// create the server
	grpcServer := grpc.NewServer()

	// register our gRPC services
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(grpcServer, orderService)

	// make the grpc server listen to tcp listener ...
	log.Println("starting grpc server on", s.addr)

	return grpcServer.Serve(lis)
}
