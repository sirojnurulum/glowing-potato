package orders

import (
	"context"
	"glowing-potato/infra/context/repository"
	"glowing-potato/objects"
)

type OrdersServiceInterface interface {
	CreateOrder(ctx context.Context, order objects.CreateOrder) error
}

func NewOrdersService(ctx *repository.RepositoryCtx) OrdersServiceInterface {
	return &ordersService{
		ctx,
	}
}
