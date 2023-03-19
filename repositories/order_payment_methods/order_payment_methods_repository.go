package order_payment_methods

import (
	"context"
	"database/sql"
	"glowing-potato/infra/db"
	"glowing-potato/objects"
)

type orderPaymentMethodsRepository struct {
	*db.DB
}

func (o orderPaymentMethodsRepository) CreateOrderPaymentMethods(ctx context.Context, tx *sql.Tx, methods objects.OrderPaymentMethods) error {
	_, err := tx.ExecContext(ctx, queryCreateOrderpaymentMethods, methods.OrderId, methods.PaymentMethodId)
	if err != nil {
		return err
	}
	return nil
}
