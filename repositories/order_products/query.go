package order_products

const (
	queryCreateOrderProduct = `
insert into order_products (order_id, product_id)
values ($1, $2);
`
)
