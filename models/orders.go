package models

type Orders struct {
	Id                int `db:"id"`
	CustomerId        int `db:"customer_id"`
	CustomerAddressId int `db:"customer_address_id"`
}

type OrderProducts struct {
	Id        int `db:"id"`
	OrderId   int `db:"order_id"`
	ProductId int `db:"product_id"`
}

type OrderPaymentMethods struct {
	Id              int `db:"id"`
	OrderId         int `db:"order_id"`
	PaymentMethodId int `db:"payment_method_id"`
}
