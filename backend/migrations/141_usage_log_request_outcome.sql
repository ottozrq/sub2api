ALTER TABLE usage_logs
    ADD COLUMN IF NOT EXISTS request_success BOOLEAN NOT NULL DEFAULT TRUE;

ALTER TABLE usage_logs
    ADD COLUMN IF NOT EXISTS error_type VARCHAR(64);

CREATE INDEX IF NOT EXISTS idx_usage_logs_request_success_created_at
    ON usage_logs (request_success, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_usage_logs_error_type_created_at
    ON usage_logs (error_type, created_at DESC)
    WHERE error_type IS NOT NULL;
