package orders

const (
	queryCreateOrders = `
insert into orders (customer_id, customer_address_id)
values ($1, $2) returning id;
`
)
