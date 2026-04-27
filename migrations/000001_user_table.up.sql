create table if not exists users (
    id serial primary key,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp null,
    email varchar(64) not null unique,
    username varchar(32) not null unique,
    firstname varchar(32) null,
    lastname varchar(32) null,
    is_superuser boolean default false,
    is_verified boolean default false
);

create index if not exists idx_users_email on users(email);
create index if not exists idx_users_username on users(username);