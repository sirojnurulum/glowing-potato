package orders

import (
	"context"
	"database/sql"
	"glowing-potato/infra/db"
	"glowing-potato/objects"
)

type OrdersRepositoryInterface interface {
	CreateOrders(ctx context.Context, tx *sql.Tx, orders objects.Orders) (*int, error)
}

func NewOrdersRepository(db *db.DB) OrdersRepositoryInterface {
	return &ordersRepository{
		db,
	}
}
