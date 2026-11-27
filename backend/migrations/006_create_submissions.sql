-- +goose Up
CREATE TABLE submissions (
    id BIGSERIAL PRIMARY KEY,
    bounty_id BIGINT NOT NULL REFERENCES bounties(id) ON DELETE CASCADE,
    executor_id BIGINT NOT NULL REFERENCES users(id),
    result TEXT NOT NULL,
    status VARCHAR(32) NOT NULL DEFAULT 'pending'
        CHECK (status IN ('pending', 'approved', 'rejected'))
);

CREATE INDEX idx_submissions_bounty_id ON submissions (bounty_id);
CREATE INDEX idx_submissions_executor_id ON submissions (executor_id);

-- +goose Down
DROP TABLE IF EXISTS submissions;
