package repository

import (
	"glowing-potato/infra/db"
	"glowing-potato/repositories/order_payment_methods"
	"glowing-potato/repositories/order_products"
	"glowing-potato/repositories/orders"
)

type RepositoryCtx struct {
	DB                            *db.DB
	OrdersRepository              orders.OrdersRepositoryInterface
	OrderProductRepository        order_products.OrderProductRepositoryInterface
	OrderPaymentMethodsRepository order_payment_methods.OrderPaymentMethodsRepositoryInterface
}
