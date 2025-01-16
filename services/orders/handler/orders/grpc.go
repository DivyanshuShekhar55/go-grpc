package handler

import (
	"github.com/DivyanshuShekhar55/go-grpc/services/common/genproto/orders"
	"github.com/DivyanshuShekhar55/go-grpc/services/orders/types"
	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	// the orderService is an interface, so basically saying our handler will have this, this and those functions
	ordersService types.OrderService
	
	// this is the fallback or default created in auto-generation code
	orders.UnimplementedOrderServiceServer
}

func NewGrpcOrdersService(grpc *grpc.Server, ordersService types.OrderService) {
	gRPCHandler := &OrdersGrpcHandler{
		ordersService: ordersService,
	}

	// register the OrderServiceServer
	// basically, we say that the handler function for orders is of type above created
	// like saying the orders handler will handle functions like creating orders, getting orders, etc
	// then we "register" this handler with the main grpc server (services/orders/grpc.go)
	// we can have another handler for another thing and also register that to the same server
	orders.RegisterOrderServiceServer(grpc, gRPCHandler)

}
