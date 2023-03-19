package orders

import (
	"glowing-potato/infra/context/service"
	"net/http"
)

type OrdersHandlerInterface interface {
	CreateOrder(w http.ResponseWriter, r *http.Request)
}

func NewOrdersHandler(svc *service.ServiceCtx) OrdersHandlerInterface {
	return &ordersHandler{
		svc,
	}
}
