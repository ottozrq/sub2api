-- Add per-subscription rolling request window quota support.

ALTER TABLE subscription_plans
    ADD COLUMN IF NOT EXISTS window_quota_count INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS window_quota_minutes INT NOT NULL DEFAULT 0;

ALTER TABLE payment_orders
    ADD COLUMN IF NOT EXISTS subscription_window_quota_count INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS subscription_window_quota_minutes INT NOT NULL DEFAULT 0;

ALTER TABLE user_subscriptions
    ADD COLUMN IF NOT EXISTS window_quota_count INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS window_quota_minutes INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS window_usage_count INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS window_start TIMESTAMPTZ;
