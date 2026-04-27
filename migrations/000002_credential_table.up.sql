create table if not exists credentials (
    id serial primary key,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at timestamp null,
    user_id int references users(id),
    type varchar(32) not null,
    key varchar(32) not null,
    value varchar(32) not null
);
create index if not exists idx_credentials_user_id_type_key on credentials(user_id, type, key);