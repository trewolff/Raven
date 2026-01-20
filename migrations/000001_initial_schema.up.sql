-- Initial schema for Raven SOAR events
CREATE TABLE IF NOT EXISTS events (
    id SERIAL PRIMARY KEY,
    key VARCHAR(255) NOT NULL,
    data JSONB,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Index for faster lookups by key
CREATE INDEX IF NOT EXISTS idx_events_key ON events(key);