package service

import (
	"context"

	"github.com/DivyanshuShekhar55/go-grpc/services/common/genproto/orders"
)

// create a mock db
var ordersDB = make([]*orders.Order, 0)

type OrderService struct {
	// store implementation
}

func NewOrderService() *OrderService {
	return &OrderService{}
}

// create the CreateOrder Business logic here which will be used
func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	ordersDB = append(ordersDB, order)
	return nil
}

// create the GetOrders Business implementation
func (s *OrderService) GetOrders(ctx context.Context) []*orders.Order {
	return ordersDB
}
