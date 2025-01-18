package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/DivyanshuShekhar55/go-grpc/services/common/genproto/orders"
)

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	return &httpServer{addr: addr}
}

func (s *httpServer) Run() error {

	// like any http server create a router
	router := http.NewServeMux()

	// create the connection to grpc server
	conn := NewGRPCClient(":9000")
	defer conn.Close()

	router.HandleFunc("/orders", func(w http.ResponseWriter, r *http.Request) {
		c := orders.NewOrderServiceClient(conn)

		// if request not approved after 3 seconds, timeout
		ctx, cancel := context.WithTimeout(r.Context(), time.Second*3)
		defer cancel()

		// this is a gRPC call, it will forward this data to the gRPC service server
		// look at the orders/handler/http.go file where the request for /orders is listened to
		_, err := c.CreateOrder(ctx, &orders.CreateOrderRequest{
			CustomerID: 24,
			ProductID:  1123,
			Quantity:   2,
		})
		if err != nil {
			log.Fatalf("client error %v", err)
		}

	})

	log.Println("starting server on :", s.addr)
	return http.ListenAndServe(s.addr, router)

}
