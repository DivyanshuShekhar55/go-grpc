package service

import "github.com/DivyanshuShekhar55/go-grpc/services/common/genproto/orders"

// create a mock db
var ordersDB = make([]*orders.Order, 0)

type OrderService struct {
	// store implementation
}

func NewOrderService() *OrderService {
	return &OrderService{}
}
