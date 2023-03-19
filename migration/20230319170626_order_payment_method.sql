-- +goose Up
-- +goose StatementBegin
create table order_payment_methods
(
    id                serial not null primary key,
    order_id          int not null,
    payment_method_id int not null,
    constraint fk_order_payment_method_orders foreign key (order_id) references orders (id),
    constraint fk_order_payment_method_methods foreign key (payment_method_id) references payment_method (id),
    constraint u_order_id_payment_method_id unique (order_id, payment_method_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
