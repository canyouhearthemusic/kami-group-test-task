-- name: GetReservation :one
SELECT * FROM reservations WHERE id = $1 LIMIT 1;

-- name: GetReservationsByRoomID :many
SELECT * FROM reservations WHERE room_id = $1;

-- name: GetAllReservations :many
SELECT * FROM reservations;

-- name: CreateReservation :one
INSERT INTO reservations(room_id, start_time, end_time) VALUES ($1, $2, $3) RETURNING *;

-- name: UpdateReservation :one
UPDATE reservations SET
    room_id = $2,
    start_time = $3,
    end_time = $4
WHERE id = $1
RETURNING *;

-- name: DeleteReservation :exec
DELETE FROM reservations WHERE id = $1;