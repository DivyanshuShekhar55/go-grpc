package handler

import (
	"net/http"

	"github.com/DivyanshuShekhar55/go-grpc/services/common/genproto/orders"
	"github.com/DivyanshuShekhar55/go-grpc/services/common/util"
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

func (h *OrdersHttpHandler) RegisterRouter(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {

	//get req thing
	var req orders.CreateOrderRequest
	//parse json
	err := util.ParseJSON(r, &req)
	if err != nil {
		util.WriteError(w, http.StatusBadRequest, err)
		return
	}

	// create a mock order ...
	// in the proto file we created a structure that had customerid, productid, quantity
	// in protofile generated we have methods to get these fields from the req, orders.pb.go line 60 ... the req is coming from kitchen/http.go
	order := &orders.Order{
		OrderID:    42,
		CustomerId: req.GetCustomerID(),
		ProductID:  req.GetProductID(),
		Quantity:   req.GetQuantity(),
	}

	// follwoing is the grpc handler, not the http one
	err = h.ordersService.CreateOrder(r.Context(), order)
	if err != nil {
		util.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	// all good, return the protobuf response and the http response
	res := &orders.CreateOrderResponse{Status: "success"}
	util.WriteJSON(w, http.StatusOK, res)

}
