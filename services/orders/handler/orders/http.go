package handler

import (
	"net/http"

	"github.com/DivyanshuShekhar55/go-grpc/services/orders/types"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewHttpOrderHandler(orderService types.OrderService) *OrdersHttpHandler {
	handler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return handler
}

func(h *OrdersHttpHandler) RegisterRouter (router *http.ServeMux){
	router.HandleFunc("POST /orders",h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request){
	
}