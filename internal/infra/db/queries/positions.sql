-- name: GetPosition :one
SELECT position_id, name, updated_at, created_at
FROM positions
WHERE position_id = $1
LIMIT 1;

-- name: ListPositions :many
SELECT position_id, name, updated_at, created_at
FROM positions
ORDER BY name;

-- name: ExistsPositionByName :one
SELECT EXISTS(
    SELECT 1 FROM positions WHERE name = $1
);

-- name: GetPositionByName :one
SELECT position_id, name, updated_at, created_at
FROM positions
WHERE name = $1
LIMIT 1;

-- name: CountPositions :one
SELECT COUNT(*) FROM positions;

-- SQL constants for mutations (used in repository)
-- name: CreatePosition :exec
INSERT INTO positions (
    position_id,
    name,
    created_at
) VALUES ($1, $2, $3);

-- name: UpdatePosition :exec
UPDATE positions
SET
    name = $2,
    updated_at = $3
WHERE position_id = $1;

-- name: DeletePosition :exec
DELETE FROM positions
WHERE position_id = $1;
