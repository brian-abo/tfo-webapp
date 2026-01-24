-- +goose Up
CREATE TABLE schema_version (
    id INTEGER PRIMARY KEY,
    version TEXT NOT NULL,
    applied_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO schema_version (id, version) VALUES (1, '1.0.0');

-- +goose Down
DROP TABLE schema_version;
