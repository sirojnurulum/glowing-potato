package handler

import "glowing-potato/cmd/http/handlers/orders"

type HandlerCtx struct {
	OrdersHandler orders.OrdersHandlerInterface
}
