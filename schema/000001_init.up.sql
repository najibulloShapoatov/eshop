CREATE TABLE users
(
    id            bigserial       not null primary key,
    account      varchar(255) not null unique,
    password     varchar(255) not null
    type         varchar(255) not null
);

CREATE TABLE products
(
    id              bigserial       not null primary key,
    name            varchar(255) not null
    description     text not null
	is_active       boolean  not null default(true)
	quantity        int4  not null default(0)
	created_at      TIMESTAMP not null default(CURRENT_TIMESTAMP)
	updated_at      TIMESTAMP not null default(CURRENT_TIMESTAMP)
);

CREATE TABLE carts
(
    id              bigserial       not null primary key,
	user_id         int8  references(users)
	created_at      TIMESTAMP not null default(CURRENT_TIMESTAMP)
	updated_at      TIMESTAMP not null default(CURRENT_TIMESTAMP)
);

CREATE TABLE cart_items
(
    id              bigserial       not null primary key,
	cart_id         int8  references(carts)
	product_id      int8  references(products)
    quantity        int4  not null default(0)
);

