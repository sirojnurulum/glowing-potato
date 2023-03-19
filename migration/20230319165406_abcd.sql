-- +goose Up
-- +goose StatementBegin
create table customer
(
    id            serial      not null primary key,
    customer_name char(50) not null
);

create table customer_address
(
    id          serial     not null primary key,
    customer_id int     not null,
    address     varchar not null,
    constraint fk_customer foreign key (customer_id) references customer (id)
);

create table product
(
    id    serial     not null primary key,
    name  varchar not null,
    price money   not null
);

create table payment_method
(
    id        serial     not null primary key,
    name      varchar not null,
    is_active bool    not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
