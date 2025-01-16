package types

import (
	"context"
	"github.com/DivyanshuShekhar55/go-grpc/services/common/genproto/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
}