CREATE TABLE IF NOT EXISTS usage_log_retry_queue (
    id BIGSERIAL PRIMARY KEY,
    request_id VARCHAR(64) NOT NULL,
    api_key_id BIGINT NOT NULL,
    log_key TEXT NOT NULL DEFAULT 'service.gateway',
    payload JSONB NOT NULL,
    status TEXT NOT NULL DEFAULT 'pending',
    attempts INT NOT NULL DEFAULT 0,
    last_error TEXT NOT NULL DEFAULT '',
    next_attempt_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    locked_until TIMESTAMPTZ,
    dead_lettered_at TIMESTAMPTZ,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE UNIQUE INDEX IF NOT EXISTS idx_usage_log_retry_queue_request_api_key
    ON usage_log_retry_queue (request_id, api_key_id);

CREATE INDEX IF NOT EXISTS idx_usage_log_retry_queue_due
    ON usage_log_retry_queue (status, next_attempt_at, locked_until, id);
