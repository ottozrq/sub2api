CREATE TABLE IF NOT EXISTS user_disposition_audits (
    id BIGSERIAL PRIMARY KEY,
    user_id BIGINT NOT NULL,
    operator_user_id BIGINT,
    reason TEXT NOT NULL DEFAULT '',
    actions JSONB NOT NULL DEFAULT '{}'::jsonb,
    summary JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

CREATE INDEX IF NOT EXISTS idx_user_disposition_audits_user_created
    ON user_disposition_audits (user_id, created_at DESC);

CREATE INDEX IF NOT EXISTS idx_user_disposition_audits_operator_created
    ON user_disposition_audits (operator_user_id, created_at DESC);
