package order_payment_methods

const (
	queryCreateOrderpaymentMethods = `
insert into order_payment_methods (order_id, payment_method_id)
values ($1, $2);
`
)
