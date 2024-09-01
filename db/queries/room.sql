-- name: GetRoom :one
SELECT * FROM rooms WHERE id = $1 LIMIT 1;

-- name: GetRooms :many
SELECT * FROM rooms;

-- name: CreateRoom :one
INSERT INTO rooms(name) VALUES ($1) RETURNING *;

-- name: UpdateRoom :one
UPDATE rooms SET name = $2 WHERE id = $1 RETURNING *;

-- name: DeleteRoom :exec
DELETE FROM rooms WHERE id = $1;