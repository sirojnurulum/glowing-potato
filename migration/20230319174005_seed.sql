-- +goose Up
-- +goose StatementBegin
insert into customer (id, customer_name)
values (1, 'abc');

insert into customer_address (customer_id, address)
values (1, 'address 1 of customer 1');
insert into customer_address (customer_id, address)
values (1, 'address 2 of customer 1');

insert into product (id, name, price)
values (1, 'product 1', 5000);
insert into product (id, name, price)
values (2, 'product 2', 6000);
insert into product (id, name, price)
values (3, 'product 3', 7000);

insert into payment_method (id, name, is_active)
values (1, 'payment method 1', true);
insert into payment_method (id, name, is_active)
values (2, 'payment method 2', true);
insert into payment_method (id, name, is_active)
values (3, 'payment method 3', true);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd