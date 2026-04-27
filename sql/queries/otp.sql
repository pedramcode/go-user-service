-- sql/queries/otp.sql

-- name: OtpGetByID :one
select * from otps
    where id = $1 and deleted_at is null
limit 1;

-- name: OtpGetValidOtp :one
select * from otps
where 
    user_id = $1 and code = $2 and 
    reason = $3 and medium = $4 and used_at is null and deleted_at is null
limit 1;

-- name: OtpDeleteByID :execrows
update otps set
    deleted_at = now()
where id = $1 and deleted_at is null;

-- name: OtpCreate :one
insert into otps (
    user_id,
    code,
    reason,
    medium,
    used_at
) values ($1, $2, $3, $4, $5)
returning id, created_at, updated_at;

-- name: OtpUpdate :one
update otps set
    user_id = $2,
    code = $3,
    reason = $4,
    medium = $5,
    used_at = $6
where id = $1 and deleted_at is null
returning updated_at;