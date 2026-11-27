-- +goose Up
CREATE TABLE join_requests (
    id BIGSERIAL PRIMARY KEY,
    room_id BIGINT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    status VARCHAR(32) NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'approved', 'declined')),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    reviewed_at TIMESTAMPTZ
);

CREATE INDEX idx_join_requests_room_id ON join_requests (room_id);
CREATE INDEX idx_join_requests_user_id ON join_requests (user_id);
CREATE UNIQUE INDEX idx_pending_join_request
    ON join_requests (room_id, user_id)
    WHERE status = 'pending';

-- +goose Down
DROP TABLE IF EXISTS join_requests;
