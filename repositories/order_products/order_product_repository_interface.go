package order_products

import (
	"context"
	"database/sql"
	"glowing-potato/infra/db"
	"glowing-potato/objects"
)

type OrderProductRepositoryInterface interface {
	CreateOrderProductRepository(ctx context.Context, tx *sql.Tx, products objects.OrderProducts) error
}

func NewOrderProductRepository(db *db.DB) OrderProductRepositoryInterface {
	return &orderProductRepository{
		db,
	}
}
