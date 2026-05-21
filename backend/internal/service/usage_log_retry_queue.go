package service

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/Wei-Shaw/sub2api/internal/pkg/logger"
)

const (
	usageLogRetryQueueSize    = 10000
	usageLogRetryWorkerCount  = 2
	usageLogRetryTimeout      = 10 * time.Second
	usageLogRetryInitialDelay = 500 * time.Millisecond
	usageLogRetryMaxDelay     = time.Minute
	usageLogRetryMaxAttempts  = 20
)

type usageLogRetryJob struct {
	log      *UsageLog
	logKey   string
	attempts int
	lastErr  string
}

type usageLogRetryQueue struct {
	repo UsageLogRepository
	ch   chan usageLogRetryJob
	once sync.Once
}

type persistentUsageLogRetryStore interface {
	EnqueueUsageLogRetry(ctx context.Context, usageLog *UsageLog, logKey string, cause error) error
}

var usageLogRetryQueues sync.Map

func enqueueUsageLogRetry(repo UsageLogRepository, usageLog *UsageLog, logKey string, cause error) {
	if repo == nil || usageLog == nil {
		return
	}
	if store, ok := repo.(persistentUsageLogRetryStore); ok {
		ctx, cancel := context.WithTimeout(context.Background(), usageLogRetryTimeout)
		err := store.EnqueueUsageLogRetry(ctx, cloneUsageLogForRetry(usageLog), logKey, cause)
		cancel()
		if err == nil {
			logger.LegacyPrintf(logKey, "Usage log persisted for retry: request_id=%s api_key_id=%d cause=%s", usageLog.RequestID, usageLog.APIKeyID, errorString(cause))
			return
		}
		logger.LegacyPrintf(logKey, "Persist usage log retry failed, falling back to memory queue: request_id=%s api_key_id=%d err=%v", usageLog.RequestID, usageLog.APIKeyID, err)
	}
	queue := getUsageLogRetryQueue(repo)
	job := usageLogRetryJob{
		log:     cloneUsageLogForRetry(usageLog),
		logKey:  logKey,
		lastErr: errorString(cause),
	}
	select {
	case queue.ch <- job:
		logger.LegacyPrintf(logKey, "Usage log queued for retry: request_id=%s api_key_id=%d cause=%s", job.log.RequestID, job.log.APIKeyID, job.lastErr)
	default:
		logger.LegacyPrintf(logKey, "Usage log retry queue full; detail may remain inconsistent: request_id=%s api_key_id=%d cause=%s", job.log.RequestID, job.log.APIKeyID, job.lastErr)
	}
}

func getUsageLogRetryQueue(repo UsageLogRepository) *usageLogRetryQueue {
	if existing, ok := usageLogRetryQueues.Load(repo); ok {
		return existing.(*usageLogRetryQueue)
	}
	queue := &usageLogRetryQueue{
		repo: repo,
		ch:   make(chan usageLogRetryJob, usageLogRetryQueueSize),
	}
	actual, _ := usageLogRetryQueues.LoadOrStore(repo, queue)
	q := actual.(*usageLogRetryQueue)
	q.start()
	return q
}

func (q *usageLogRetryQueue) start() {
	q.once.Do(func() {
		for i := 0; i < usageLogRetryWorkerCount; i++ {
			go q.worker()
		}
	})
}

func (q *usageLogRetryQueue) worker() {
	for job := range q.ch {
		q.retryUntilPersisted(job)
	}
}

func (q *usageLogRetryQueue) retryUntilPersisted(job usageLogRetryJob) {
	delay := usageLogRetryInitialDelay
	for {
		if job.attempts > 0 {
			time.Sleep(delay)
			delay *= 2
			if delay > usageLogRetryMaxDelay {
				delay = usageLogRetryMaxDelay
			}
		}
		job.attempts++

		ctx, cancel := context.WithTimeout(context.Background(), usageLogRetryTimeout)
		inserted, err := q.repo.Create(ctx, cloneUsageLogForRetry(job.log))
		cancel()
		if err == nil {
			if inserted {
				logger.LegacyPrintf(job.logKey, "Usage log retry persisted: request_id=%s api_key_id=%d attempts=%d", job.log.RequestID, job.log.APIKeyID, job.attempts)
			} else {
				logger.LegacyPrintf(job.logKey, "Usage log retry skipped duplicate: request_id=%s api_key_id=%d attempts=%d", job.log.RequestID, job.log.APIKeyID, job.attempts)
			}
			return
		}
		logger.LegacyPrintf(job.logKey, "Usage log retry failed: request_id=%s api_key_id=%d attempts=%d err=%v", job.log.RequestID, job.log.APIKeyID, job.attempts, err)
		if job.attempts >= usageLogRetryMaxAttempts {
			logger.LegacyPrintf(job.logKey, "Usage log retry abandoned after max attempts: request_id=%s api_key_id=%d attempts=%d last_err=%v", job.log.RequestID, job.log.APIKeyID, job.attempts, err)
			return
		}
	}
}

func cloneUsageLogForRetry(in *UsageLog) *UsageLog {
	if in == nil {
		return nil
	}
	out := *in
	out.User = nil
	out.APIKey = nil
	out.Account = nil
	out.Group = nil
	out.Subscription = nil
	out.UpstreamModel = cloneUsageLogStringPtr(in.UpstreamModel)
	out.ChannelID = cloneUsageLogInt64Ptr(in.ChannelID)
	out.ModelMappingChain = cloneUsageLogStringPtr(in.ModelMappingChain)
	out.BillingTier = cloneUsageLogStringPtr(in.BillingTier)
	out.BillingMode = cloneUsageLogStringPtr(in.BillingMode)
	out.ServiceTier = cloneUsageLogStringPtr(in.ServiceTier)
	out.ReasoningEffort = cloneUsageLogStringPtr(in.ReasoningEffort)
	out.InboundEndpoint = cloneUsageLogStringPtr(in.InboundEndpoint)
	out.UpstreamEndpoint = cloneUsageLogStringPtr(in.UpstreamEndpoint)
	out.GroupID = cloneUsageLogInt64Ptr(in.GroupID)
	out.SubscriptionID = cloneUsageLogInt64Ptr(in.SubscriptionID)
	out.AccountRateMultiplier = cloneUsageLogFloat64Ptr(in.AccountRateMultiplier)
	out.AccountStatsCost = cloneUsageLogFloat64Ptr(in.AccountStatsCost)
	out.DurationMs = cloneUsageLogIntPtr(in.DurationMs)
	out.FirstTokenMs = cloneUsageLogIntPtr(in.FirstTokenMs)
	out.UserAgent = cloneUsageLogStringPtr(in.UserAgent)
	out.IPAddress = cloneUsageLogStringPtr(in.IPAddress)
	out.ImageSize = cloneUsageLogStringPtr(in.ImageSize)
	out.MediaType = cloneUsageLogStringPtr(in.MediaType)
	return &out
}

func cloneUsageLogStringPtr(in *string) *string {
	if in == nil {
		return nil
	}
	out := *in
	return &out
}

func cloneUsageLogInt64Ptr(in *int64) *int64 {
	if in == nil {
		return nil
	}
	out := *in
	return &out
}

func cloneUsageLogIntPtr(in *int) *int {
	if in == nil {
		return nil
	}
	out := *in
	return &out
}

func cloneUsageLogFloat64Ptr(in *float64) *float64 {
	if in == nil {
		return nil
	}
	out := *in
	return &out
}

func errorString(err error) string {
	if err == nil {
		return ""
	}
	return fmt.Sprintf("%v", err)
}
