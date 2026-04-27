-- sql/queries/credential.sql

-- name: CredentialGetByID :one
select * from credentials
    where id = $1 and deleted_at is null
limit 1;

-- name: CredentialGetByUserTypeKey :one
select * from credentials
    where user_id = $1 and type = $2 and key = $3 and deleted_at is null
limit 1;

-- name: CredentialCreate :one
insert into credentials (
    user_id, type, key, value
)
values
(
    $1, $2, $3, $4
)
returning id, created_at, updated_at;

-- name: CredentialUpdate :one
update credentials set
    user_id = $2,
    type = $3,
    key = $4,
    value = $5
where id = $1 and deleted_at is null
returning updated_at;

-- name: CredentialDeleteByID :execrows
update credentials set
    deleted_at = now()
where id = $1 and deleted_at is null;