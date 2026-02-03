-- name: GetStaff :one
SELECT
    staff_id,
    position_id,
    first_name,
    last_name,
    email,
    password_hash,
    password_salt,
    status,
    selfie_url,
    selfie_thumb_url,
    updated_at,
    created_at
FROM staffs
WHERE staff_id = $1
LIMIT 1;

-- name: GetStaffByEmail :one
SELECT
    staff_id,
    position_id,
    first_name,
    last_name,
    email,
    password_hash,
    password_salt,
    status,
    selfie_url,
    selfie_thumb_url,
    updated_at,
    created_at
FROM staffs
WHERE email = $1
LIMIT 1;

-- name: ListStaffs :many
SELECT
    staff_id,
    position_id,
    first_name,
    last_name,
    email,
    password_hash,
    password_salt,
    status,
    selfie_url,
    selfie_thumb_url,
    updated_at,
    created_at
FROM staffs
ORDER BY created_at DESC;

-- name: ListStaffsByPosition :many
SELECT
    staff_id,
    position_id,
    first_name,
    last_name,
    email,
    password_hash,
    password_salt,
    status,
    selfie_url,
    selfie_thumb_url,
    updated_at,
    created_at
FROM staffs
WHERE position_id = $1
ORDER BY created_at DESC;

-- name: ExistsStaffByEmail :one
SELECT EXISTS(
    SELECT 1 FROM staffs WHERE email = $1
);

-- name: CountStaffs :one
SELECT COUNT(*) FROM staffs;

-- name: CountStaffsByStatus :one
SELECT COUNT(*) FROM staffs WHERE status = $1;

-- SQL constants for mutations (used in repository)
-- name: CreateStaff :exec
INSERT INTO staffs (
    staff_id,
    position_id,
    first_name,
    last_name,
    email,
    password_hash,
    password_salt,
    status,
    selfie_url,
    selfie_thumb_url,
    created_at
) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11);

-- name: UpdateStaff :exec
UPDATE staffs
SET
    first_name = $2,
    last_name = $3,
    email = $4,
    password_hash = $5,
    password_salt = $6,
    status = $7,
    selfie_url = $8,
    selfie_thumb_url = $9,
    updated_at = $10
WHERE staff_id = $1;

-- name: UpdateStaffStatus :exec
UPDATE staffs
SET
    status = $2,
    updated_at = $3
WHERE staff_id = $1;

-- name: DeleteStaff :exec
DELETE FROM staffs
WHERE staff_id = $1;
