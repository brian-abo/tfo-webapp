-- +goose Up
CREATE TABLE signups (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id UUID NOT NULL REFERENCES users(id),
    hunt_id UUID NOT NULL REFERENCES hunts(id),
    eligibility_snapshot TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    withdrawn_at TIMESTAMPTZ,

    CONSTRAINT signups_user_hunt_unique UNIQUE (user_id, hunt_id)
);

CREATE INDEX idx_signups_user_id ON signups (user_id);
CREATE INDEX idx_signups_hunt_id ON signups (hunt_id);

-- +goose Down
DROP TABLE signups;
