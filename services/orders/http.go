package main

import (
	"log"
	"net/http"

	handler "github.com/DivyanshuShekhar55/go-grpc/services/orders/handler/orders"
	"github.com/DivyanshuShekhar55/go-grpc/services/orders/service"
)

// this is the place where we can handle the http services / routes ... just fopr example purposes
// we are using just the net/http lib but you might want to use chi (or smthing else) too

type httpServer struct {
	addr string
}

func NewHttpServer(addr string) *httpServer {
	srv := &httpServer{addr: addr}
	return srv
}

func (s *httpServer) Run() error {

	// create a router
	r := http.NewServeMux()

	// register all the http routing services
	// the NewOrderService will ideally be a interface of func handlers
	orderService := service.NewOrderService()
	orderHandler := handler.NewHttpOrderHandler(orderService)

	// following will add all routes of http to this main mux 'r'
	orderHandler.RegisterRouter(r)

	log.Println("Starting Server On", s.addr)
	return http.ListenAndServe(s.addr, r)
}
