CREATE TABLE IF NOT EXISTS users
(
    id serial not null unique,
    first_name varchar(255) unique,
    last_name varchar(255) unique
);