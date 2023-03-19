package order_products

import (
	"context"
	"database/sql"
	"glowing-potato/infra/db"
	"glowing-potato/objects"
)

type orderProductRepository struct {
	*db.DB
}

func (o orderProductRepository) CreateOrderProductRepository(ctx context.Context, tx *sql.Tx, products objects.OrderProducts) error {
	_, err := tx.ExecContext(ctx, queryCreateOrderProduct, products.OrderId, products.ProductId)
	if err != nil {
		return err
	}
	return nil
}
