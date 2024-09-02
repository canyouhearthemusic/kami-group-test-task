CREATE TABLE IF NOT EXISTS reservations(
    id SERIAL PRIMARY KEY,
    room_id VARCHAR(10) NOT NULL,
    start_time TIMESTAMP NOT NULL,
    end_time TIMESTAMP NOT NULL
);

CREATE INDEX idx_reservations_room_id ON reservations(room_id);