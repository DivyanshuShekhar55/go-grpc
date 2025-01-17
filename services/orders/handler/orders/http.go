package handler

import "github.com/DivyanshuShekhar55/go-grpc/services/orders/types"

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrderHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return handler
}
