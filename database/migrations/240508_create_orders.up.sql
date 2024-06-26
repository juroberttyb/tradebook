CREATE TABLE IF NOT EXISTS public."order"
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    action character varying(8) COLLATE pg_catalog."default" NOT NULL,
    price integer NOT NULL,
    quantity integer NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT now(),
    CONSTRAINT orders_pkey PRIMARY KEY (id)
)