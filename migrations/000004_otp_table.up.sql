create table if not exists otps (
    id serial primary key,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp null,
    user_id int references users(id) not null,
    reason varchar(32) not null,
    medium varchar(32) not null,
    code char(8) not null,
    used_at timestamp null
);