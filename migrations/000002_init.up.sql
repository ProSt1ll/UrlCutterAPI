-- +goose Up
CREATE TABLE urls
(
    id        serial       not null unique,
    long_url  varchar(255) not null unique,
    short_url varchar(255) not null unique
);

CREATE TABLE config(
    id        serial       not null unique,
    url_cnt int,
    list varchar(255) not null
);




-- +goose Down
DROP TABLE URLs;

DROP TABLE config;
