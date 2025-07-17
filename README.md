First Copy File `.env_example` into folder `cmd`

Rename `.env_example` into `.env`

Edit `CONNECTION_STRING` at `.env` File with your `connection string postgreSQL`

Here's DDL you can copy on your database
```
CREATE TABLE public.product (
	id_product uuid DEFAULT uuid_in(md5(random()::text || random()::text)::cstring) NOT NULL,
	product_name varchar NULL,
	price int4 NULL,
	stock int4 NULL,
	created_at timestamp NULL,
	updated_at timestamp NULL,
	deleted_at timestamp NULL,
	CONSTRAINT product_pk PRIMARY KEY (id_product)
);
```
