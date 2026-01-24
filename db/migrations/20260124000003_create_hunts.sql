-- +goose Up
CREATE TABLE hunts (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    location TEXT NOT NULL,
    image_urls JSONB NOT NULL DEFAULT '[]',
    qualifiers TEXT,
    hunt_date TIMESTAMPTZ NOT NULL,
    signup_window_start TIMESTAMPTZ NOT NULL,
    signup_window_end TIMESTAMPTZ NOT NULL,
    primary_capacity INTEGER NOT NULL,
    alternate_capacity INTEGER NOT NULL,
    status TEXT NOT NULL DEFAULT 'draft',
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),

    CONSTRAINT hunts_status_check CHECK (
        status IN ('draft', 'open', 'closed', 'completed', 'cancelled')
    ),
    CONSTRAINT hunts_capacity_positive CHECK (primary_capacity >= 0 AND alternate_capacity >= 0)
);

CREATE INDEX idx_hunts_status ON hunts (status);
CREATE INDEX idx_hunts_hunt_date ON hunts (hunt_date);

-- +goose Down
DROP TABLE hunts;
