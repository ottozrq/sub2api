-- Add quota-pack support for fixed-call-count products.

ALTER TABLE subscription_plans
    ADD COLUMN IF NOT EXISTS plan_type VARCHAR(20) NOT NULL DEFAULT 'subscription',
    ADD COLUMN IF NOT EXISTS quota_count INTEGER NOT NULL DEFAULT 0;

ALTER TABLE payment_orders
    ADD COLUMN IF NOT EXISTS subscription_plan_type VARCHAR(20) NOT NULL DEFAULT 'subscription',
    ADD COLUMN IF NOT EXISTS subscription_quota_count INTEGER NOT NULL DEFAULT 0;

ALTER TABLE user_subscriptions
    ADD COLUMN IF NOT EXISTS quota_total_count INTEGER NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS quota_used_count INTEGER NOT NULL DEFAULT 0;

CREATE INDEX IF NOT EXISTS idx_subscription_plans_plan_type ON subscription_plans(plan_type);
