DROP TABLE IF EXISTS public.url CASCADE;

CREATE TABLE IF NOT EXISTS public.urls
(
    url_id SERIAL PRIMARY KEY ,
    short_url varchar(256) NOT NULL ,
    long_url varchar(256) UNIQUE NOT NULL,
    user_id int
);

CREATE TABLE IF NOT EXISTS public.users
(
    user_id SERIAL PRIMARY KEY ,
    login   VARCHAR(256) UNIQUE not null ,
    password varchar(256) not null
);

ALTER TABLE public.urls DROP CONSTRAINT IF EXISTS fk_user_id;
ALTER TABLE public.urls ADD CONSTRAINT fk_user_id FOREIGN KEY (user_id)
    REFERENCES public.users (user_id) ON DELETE CASCADE;