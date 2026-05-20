package service

import (
	"time"
)

// SubscriptionCacheData represents cached subscription data
type SubscriptionCacheData struct {
	Status             string
	ExpiresAt          time.Time
	DailyUsage         float64
	WeeklyUsage        float64
	MonthlyUsage       float64
	WindowQuotaCount   int
	WindowQuotaMinutes int
	WindowUsageCount   int
	WindowStart        *time.Time
	QuotaTotalCount    int
	QuotaUsedCount     int
	Version            int64
}
