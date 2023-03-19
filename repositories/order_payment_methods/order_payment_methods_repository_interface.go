package order_payment_methods

import (
	"context"
	"database/sql"
	"glowing-potato/infra/db"
	"glowing-potato/objects"
)

type OrderPaymentMethodsRepositoryInterface interface {
	CreateOrderPaymentMethods(ctx context.Context, tx *sql.Tx, methods objects.OrderPaymentMethods) error
}

func NewOrderPaymentMethodsRepository(db *db.DB) OrderPaymentMethodsRepositoryInterface {
	return &orderPaymentMethodsRepository{
		db,
	}
}
