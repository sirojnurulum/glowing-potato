package orders

import (
	"context"
	"glowing-potato/infra/context/repository"
	"glowing-potato/objects"
)

type ordersService struct {
	*repository.RepositoryCtx
}

func (o ordersService) CreateOrder(ctx context.Context, order objects.CreateOrder) error {
	tx, err := o.DB.Master().BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	//create order
	id, err := o.OrdersRepository.CreateOrders(ctx, tx, objects.Orders{
		CustomerId:        order.CustomerId,
		CustomerAddressId: order.AddressId,
	})
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	//create order products
	for _, x := range order.ProductIds {
		err = o.OrderProductRepository.CreateOrderProductRepository(ctx, tx, objects.OrderProducts{
			OrderId:   *id,
			ProductId: x,
		})
		if err != nil {
			_ = tx.Rollback()
			return err
		}
	}

	//create order payment method
	err = o.OrderPaymentMethodsRepository.CreateOrderPaymentMethods(ctx, tx, objects.OrderPaymentMethods{
		OrderId:         *id,
		PaymentMethodId: order.PaymentMethodId,
	})
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return nil
}
