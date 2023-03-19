package orders

import (
	"context"
	"database/sql"
	"glowing-potato/infra/db"
	"glowing-potato/objects"
)

type ordersRepository struct {
	*db.DB
}

func (o ordersRepository) CreateOrders(ctx context.Context, tx *sql.Tx, orders objects.Orders) (*int, error) {
	var id int
	err := tx.QueryRowContext(ctx, queryCreateOrders, orders.CustomerId, orders.CustomerAddressId).Scan(&id)
	if err != nil {
		return nil, err
	}
	return &id, nil
}
