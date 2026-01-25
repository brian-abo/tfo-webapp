-- +goose Up
CREATE TABLE lottery_results (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hunt_id UUID NOT NULL REFERENCES hunts(id),
    signup_id UUID NOT NULL REFERENCES signups(id),
    position INTEGER NOT NULL,
    audit_seed BIGINT NOT NULL,
    algorithm_version TEXT NOT NULL,
    drawn_at TIMESTAMPTZ NOT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT lottery_results_signup_unique UNIQUE (signup_id),
    CONSTRAINT lottery_results_position_positive CHECK (position > 0)
);

CREATE INDEX idx_lottery_results_hunt_id ON lottery_results (hunt_id);
CREATE INDEX idx_lottery_results_position ON lottery_results (hunt_id, position);

-- +goose Down
DROP TABLE lottery_results;
