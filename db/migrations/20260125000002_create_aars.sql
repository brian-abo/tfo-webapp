-- +goose Up
CREATE TABLE hunt_after_action_reports (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    hunt_id UUID NOT NULL REFERENCES hunts(id),
    description TEXT NOT NULL,
    image_urls JSONB NOT NULL DEFAULT '[]',
    created_by_id UUID NOT NULL REFERENCES users(id),
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT aars_hunt_unique UNIQUE (hunt_id)
);

CREATE INDEX idx_aars_hunt_id ON hunt_after_action_reports (hunt_id);
CREATE INDEX idx_aars_created_by_id ON hunt_after_action_reports (created_by_id);

CREATE TABLE aar_participants (
    aar_id UUID NOT NULL REFERENCES hunt_after_action_reports(id) ON DELETE CASCADE,
    user_id UUID NOT NULL REFERENCES users(id),

    PRIMARY KEY (aar_id, user_id)
);

CREATE INDEX idx_aar_participants_user_id ON aar_participants (user_id);

-- +goose Down
DROP TABLE aar_participants;
DROP TABLE hunt_after_action_reports;
