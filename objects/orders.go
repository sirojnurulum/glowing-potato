package objects

type Orders struct {
	Id                int `json:"id"`
	CustomerId        int `json:"customer_id"`
	CustomerAddressId int `json:"customer_address_id"`
}

type OrderProducts struct {
	Id        int `json:"id"`
	OrderId   int `json:"order_id"`
	ProductId int `json:"product_id"`
}

type OrderPaymentMethods struct {
	Id              int `json:"id"`
	OrderId         int `json:"order_id"`
	PaymentMethodId int `json:"payment_method_id"`
}

type CreateOrder struct {
	CustomerId      int   `json:"customer_id"`
	AddressId       int   `json:"address_id"`
	ProductIds      []int `json:"product_ids"`
	PaymentMethodId int   `json:"payment_method_id"`
}
