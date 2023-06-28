CREATE TABLE IF NOT EXISTS "users" (
    id serial primary key,
    name varchar(255) unique,
    email varchar(255) unique,
    password varchar(255) not null,
    email_verified_at timestamp,
    created_at timestamp default now()
);