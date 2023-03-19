-- +goose Up
-- +goose StatementBegin
create table orders
(
    id                  serial not null primary key,
    customer_id         int not null,
    customer_address_id int not null,
    constraint fk_order_customer foreign key (customer_id) references customer (id),
    constraint fk_order_customer_address foreign key (customer_address_id) references customer_address (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
