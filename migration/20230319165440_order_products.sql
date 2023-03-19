-- +goose Up
-- +goose StatementBegin
create table order_products
(
    id         serial not null primary key,
    order_id   int not null,
    product_id int not null,
    constraint fk_order_product_order foreign key (order_id) references orders (id),
    constraint fk_order_product_product foreign key (product_id) references product (id),
    constraint u_order_id_product_id unique (order_id, product_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
