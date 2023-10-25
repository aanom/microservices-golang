create extension if not exists "uuid-ossp";

CREATE TABLE public.users
(
    id uuid DEFAULT uuid_generate_v4(),
    username character varying NOT NULL,
    email character varying NOT NULL,
    password text NOT NULL,
    mobile character varying NOT NULL,
    address text NOT NULL,
    pincode character varying NOT NULL,
    status boolean DEFAULT true,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone,
    PRIMARY KEY (id)
);

ALTER TABLE public.users
    OWNER to postgres;