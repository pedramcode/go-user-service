-- sql/queries/users.sql

-- name: UserGetByID :one
SELECT * from users
    where id = $1 and deleted_at is null
limit 1;

-- name: UserGetByEmail :one
SELECT * from users
    where email = $1 and deleted_at is null
limit 1;

-- name: UserGetByUsername :one
SELECT * from users
    where username = $1 and deleted_at is null
limit 1;

-- name: UserCreate :one
insert into users (
    email, username, firstname,
    lastname, is_superuser, is_verified
) values 
($1, $2, $3, $4, $5, $6)
returning id, created_at, updated_at;

-- name: UserUpdate :one
update users set
    updated_at = default,
    email = $2,
    username = $3,
    firstname = $4,
    lastname = $5,
    is_superuser = $6,
    is_verified = $7
where id = $1 and deleted_at is null
returning updated_at;

-- name: UserDeleteByID :execrows
update users set
    deleted_at = now()
where id = $1 and deleted_at is null;

-- name: UserDeleteByEmail :execrows
update users set
    deleted_at = now()
where email = $1 and deleted_at is null;

-- name: UserDeleteByUsername :execrows
update users set
    deleted_at = now()
where username = $1 and deleted_at is null;

-- name: UserExistsByEmail :one
select exists (
    select 1 from users
    where email = $1 and deleted_at is null
) as exists;

-- name: UserExistsByUsername :one
select exists (
    select 1 from users
    where username = $1 and deleted_at is null
) as exists;