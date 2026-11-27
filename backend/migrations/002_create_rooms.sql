-- +goose Up
CREATE TABLE rooms (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR(128) NOT NULL,
    invite_code VARCHAR(64) NOT NULL,
    master_id BIGINT NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX idx_rooms_invite_code ON rooms (invite_code);
CREATE INDEX idx_rooms_master_id ON rooms (master_id);

-- +goose Down
DROP TABLE IF EXISTS rooms;
