-- +goose Up
create table "user" (
    id serial primary key,
    "name" text not null,
    email text not null,
    "role" int not null,
    "password" text not null,
    "password_confirm" text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);

-- +goose Down
drop table "user";