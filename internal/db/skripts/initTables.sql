-- DROP TABLE IF EXISTS public.url CASCADE;

CREATE TABLE IF NOT EXISTS public.url
(
    url_id SERIAL PRIMARY KEY ,
    short_url varchar(256) NOT NULL ,
    long_url varchar(256) UNIQUE NOT NULL
);