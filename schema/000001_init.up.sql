CREATE TABLE users if not EXISTS
(
    id            bigserial       not null primary key,
    account      varchar(255) not null unique,
    password     varchar(255) not null
    type         varchar(255) not null
);

CREATE TABLE products if not EXISTS
(
    id              bigserial       not null primary key,
    name            varchar(255) not null
    description     text not null
	is_active       boolean  not null default(true)
	quantity        int4  not null default(0)
	created_at      TIMESTAMP not null default(CURRENT_TIMESTAMP)
	updated_at      TIMESTAMP not null default(CURRENT_TIMESTAMP)
);

CREATE TABLE carts if not EXISTS
(
    id              bigserial       not null primary key,
	user_id         int8  references(users)
	created_at      TIMESTAMP not null default(CURRENT_TIMESTAMP)
	updated_at      TIMESTAMP not null default(CURRENT_TIMESTAMP)
);

CREATE TABLE cart_items if not EXISTS
(
    id              bigserial       not null primary key,
	cart_id         int8  references(carts)
	product_id      int8  references(products)
    quantity        int4  not null default(0)
);

INSERT INTO users(account, password, type) VALUES ( 'test', '75575452697d374546433739627a657e4b7a2b5362336e6340343175435724759f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08', 'manager');
