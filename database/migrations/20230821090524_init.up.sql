-- This script was generated by the ERD tool in pgAdmin 4.
-- Please log an issue at https://redmine.postgresql.org/projects/pgadmin4/issues/new if you find any bugs, including reproduction steps.
BEGIN;

CREATE EXTENSION IF NOT EXISTS hstore;

CREATE TABLE IF NOT EXISTS public.app
(
    id uuid NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    bundle_id character varying(100) COLLATE pg_catalog."default" NOT NULL,
    key_name character varying(200) COLLATE pg_catalog."default",
    seq smallint NOT NULL DEFAULT 0,
    CONSTRAINT app_pkey PRIMARY KEY (id),
    CONSTRAINT app_bundle_id_key UNIQUE (bundle_id)
);

CREATE TABLE IF NOT EXISTS public."user"
(
    id uuid NOT NULL,
    app_id uuid NOT NULL,
    external_user_id uuid,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    picture character varying(100) COLLATE pg_catalog."default",
    push_token character varying(200) COLLATE pg_catalog."default",
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    attrs hstore NOT NULL DEFAULT ''::hstore,
    email character varying(320) COLLATE pg_catalog."default",
    password_bcrypt character varying(60) COLLATE pg_catalog."default",
    type smallint NOT NULL DEFAULT 0,
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_app_id_external_user_id_key UNIQUE (app_id, external_user_id),
    CONSTRAINT user_email_key UNIQUE (email)
);

CREATE TABLE IF NOT EXISTS public.campaign
(
    id uuid NOT NULL,
    name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    app_id uuid NOT NULL,
    schedule timestamp without time zone NOT NULL,
    moderator_id uuid NOT NULL,
    audience_size bigint,
    filter jsonb,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    broadcast_message_ids uuid[],
    status smallint NOT NULL DEFAULT 0,
    CONSTRAINT campaign_pkey PRIMARY KEY (id),
    CONSTRAINT campaign_moderator_id_fkey FOREIGN KEY (moderator_id)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

CREATE TABLE IF NOT EXISTS public.message
(
    id uuid NOT NULL,
    type smallint NOT NULL,
    body text COLLATE pg_catalog."default",
    chat_id uuid NOT NULL,
    sender_id uuid NOT NULL,
    created_at timestamp without time zone NOT NULL,
    reply_to_message_id uuid,
    status smallint NOT NULL DEFAULT 0,
    media_ids uuid[],
    is_broadcast boolean NOT NULL,
    broadcast_message_id uuid,
    CONSTRAINT message_pkey PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.broadcast_message
(
    id uuid NOT NULL,
    campaign_id uuid NOT NULL,
    sender_id uuid NOT NULL,
    type smallint NOT NULL,
    body text COLLATE pg_catalog."default",
    media_ids uuid[],
    CONSTRAINT broadcast_message_pkey PRIMARY KEY (id),
    CONSTRAINT broadcast_message_sender_id_fkey FOREIGN KEY (sender_id)
        REFERENCES public."user" (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
        NOT VALID
);

CREATE TABLE IF NOT EXISTS public.chat
(
    id uuid NOT NULL,
    updated_at timestamp without time zone NOT NULL,
    last_message_id uuid,
    PRIMARY KEY (id)
);

CREATE TABLE IF NOT EXISTS public.chat_thread
(
    chat_id uuid NOT NULL,
    sender_id uuid NOT NULL,
    receiver_id uuid NOT NULL,
    unread_count bigint NOT NULL DEFAULT 0,
    last_seen_at timestamp without time zone,
    status smallint NOT NULL DEFAULT 0,
    CONSTRAINT chat_thread_pk PRIMARY KEY (sender_id, receiver_id)
);

CREATE TABLE IF NOT EXISTS public.analytics
(
    broadcast_message_id uuid NOT NULL,
    total_impressions bigint NOT NULL DEFAULT 0,
    unique_impressions bigint NOT NULL DEFAULT 0,
    total_clicks bigint NOT NULL DEFAULT 0,
    unique_clicks bigint NOT NULL DEFAULT 0,
    PRIMARY KEY (broadcast_message_id)
);

CREATE TABLE IF NOT EXISTS public.media
(
    id uuid NOT NULL,
    url character varying(100) NOT NULL,
    preview_url character varying(100),
    placeholder character varying(50),
    type smallint NOT NULL DEFAULT 0,
    PRIMARY KEY (id)
);


ALTER TABLE IF EXISTS public.message
    ADD FOREIGN KEY (chat_id)
    REFERENCES public.chat (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.message
    ADD FOREIGN KEY (sender_id)
    REFERENCES public."user" (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.message
    ADD FOREIGN KEY (broadcast_message_id)
    REFERENCES public.broadcast_message (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public."user"
    ADD CONSTRAINT user_app_id_fkey FOREIGN KEY (app_id)
    REFERENCES public.app (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.chat_thread
    ADD FOREIGN KEY (chat_id)
    REFERENCES public.chat (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.chat_thread
    ADD FOREIGN KEY (sender_id)
    REFERENCES public."user" (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.chat_thread
    ADD FOREIGN KEY (receiver_id)
    REFERENCES public."user" (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS public.analytics
    ADD FOREIGN KEY (broadcast_message_id)
    REFERENCES public.broadcast_message (id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

END;
