-- Move rolling request window quota configuration to subscription groups.

ALTER TABLE groups
    ADD COLUMN IF NOT EXISTS window_quota_count INT NOT NULL DEFAULT 0,
    ADD COLUMN IF NOT EXISTS window_quota_minutes INT NOT NULL DEFAULT 0;

UPDATE groups g
SET
    window_quota_count = p.window_quota_count,
    window_quota_minutes = p.window_quota_minutes
FROM (
    SELECT DISTINCT ON (group_id)
        group_id,
        window_quota_count,
        window_quota_minutes
    FROM subscription_plans
    WHERE window_quota_count > 0
      AND window_quota_minutes > 0
    ORDER BY group_id, sort_order ASC, id ASC
) p
WHERE g.id = p.group_id
  AND g.window_quota_count = 0
  AND g.window_quota_minutes = 0;
