-- +goose Up
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    email TEXT NOT NULL,
    name TEXT NOT NULL,
    phone TEXT,
    branch_of_service TEXT NOT NULL,
    role TEXT NOT NULL DEFAULT 'member',
    membership_status TEXT NOT NULL DEFAULT 'pending',
    facebook_id TEXT,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    deleted_at TIMESTAMPTZ,

    CONSTRAINT users_email_unique UNIQUE (email),
    CONSTRAINT users_facebook_id_unique UNIQUE (facebook_id),
    CONSTRAINT users_role_check CHECK (role IN ('member', 'staff', 'admin')),
    CONSTRAINT users_membership_status_check CHECK (
        membership_status IN ('pending', 'active', 'inactive', 'suspended')
    )
);

CREATE INDEX idx_users_active ON users (deleted_at) WHERE deleted_at IS NULL;
CREATE INDEX idx_users_email ON users (email) WHERE deleted_at IS NULL;

-- +goose Down
DROP TABLE users;
