-- +goose Up
CREATE TABLE points_transactions (
    id BIGSERIAL PRIMARY KEY,
    room_id BIGINT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    user_id BIGINT NOT NULL REFERENCES users(id),
    bounty_id BIGINT NOT NULL REFERENCES bounties(id),
    amount INTEGER NOT NULL CHECK (amount > 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    CONSTRAINT uq_points_transactions_bounty UNIQUE (bounty_id)
);

CREATE INDEX idx_points_transactions_room_id ON points_transactions (room_id);
CREATE INDEX idx_points_transactions_user_id ON points_transactions (user_id);

-- +goose Down
DROP TABLE IF EXISTS points_transactions;
