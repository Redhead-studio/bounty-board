-- +goose Up
CREATE TABLE room_members (
    id BIGSERIAL PRIMARY KEY,
    room_id BIGINT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    role VARCHAR(32) NOT NULL DEFAULT 'member'
        CHECK (role IN ('member', 'master')),
    status VARCHAR(32) NOT NULL DEFAULT 'active'
        CHECK (status IN ('active', 'blocked')),
    joined_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_room_members_room_user UNIQUE (room_id, user_id)
);

CREATE INDEX idx_room_members_room_id ON room_members (room_id);
CREATE INDEX idx_room_members_user_id ON room_members (user_id);

-- +goose Down
DROP TABLE IF EXISTS room_members;
