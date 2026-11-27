-- +goose Up
CREATE TABLE bounties (
    id BIGSERIAL PRIMARY KEY,
    room_id BIGINT NOT NULL REFERENCES rooms(id) ON DELETE CASCADE,
    author_id BIGINT NOT NULL REFERENCES users(id),
    executor_id BIGINT REFERENCES users(id),
    title VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    reward INTEGER NOT NULL CHECK (reward > 0),
    status VARCHAR(32) NOT NULL DEFAULT 'open'
        CHECK (status IN ('open', 'in_progress', 'review', 'completed', 'rejected')),
    CONSTRAINT chk_bounty_author_not_executor
        CHECK (executor_id IS NULL OR executor_id <> author_id)
);

CREATE INDEX idx_bounties_room_id ON bounties (room_id);
CREATE INDEX idx_bounties_author_id ON bounties (author_id);
CREATE INDEX idx_bounties_executor_id ON bounties (executor_id);
CREATE INDEX idx_bounties_status ON bounties (status);

-- +goose Down
DROP TABLE IF EXISTS bounties;
