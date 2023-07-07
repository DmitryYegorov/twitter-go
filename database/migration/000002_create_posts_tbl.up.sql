CREATE TABLE IF NOT EXISTS "posts" (
    id serial primary key,
    text varchar(500),
    created_by int references "users"(id),
    created_at timestamp default now()
)