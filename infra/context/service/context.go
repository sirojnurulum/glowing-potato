package service

import "glowing-potato/services/orders"

type ServiceCtx struct {
	OrdersService orders.OrdersServiceInterface
}
