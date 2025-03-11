CREATE TABLE urls 
(
    id serial primary key,
    long_url text not null,
    short_url varchar(10),
    created_at timestamp default now()
);

CREATE UNIQUE INDEX idx_short_url ON urls(short_url);