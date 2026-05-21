package service

import (
	"context"
	"errors"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type usageLogRetryRepoStub struct {
	UsageLogRepository
	failuresRemaining atomic.Int32
	calls             atomic.Int32
	done              chan struct{}
	lastLog           atomic.Pointer[UsageLog]
}

func (s *usageLogRetryRepoStub) Create(ctx context.Context, log *UsageLog) (bool, error) {
	s.calls.Add(1)
	s.lastLog.Store(log)
	if s.failuresRemaining.Add(-1) >= 0 {
		return false, errors.New("temporary insert failure")
	}
	select {
	case <-s.done:
	default:
		close(s.done)
	}
	return true, nil
}

func TestUsageLogRetryQueueRetriesUntilPersisted(t *testing.T) {
	repo := &usageLogRetryRepoStub{done: make(chan struct{})}
	repo.failuresRemaining.Store(2)
	upstreamModel := "mapped-model"
	usageLog := &UsageLog{
		UserID:        1,
		APIKeyID:      2,
		AccountID:     3,
		RequestID:     "retry-request",
		Model:         "requested-model",
		UpstreamModel: &upstreamModel,
		User:          &User{ID: 1},
		APIKey:        &APIKey{ID: 2},
		Account:       &Account{ID: 3},
	}

	enqueueUsageLogRetry(repo, usageLog, "service.test", errors.New("initial insert failure"))

	select {
	case <-repo.done:
	case <-time.After(5 * time.Second):
		t.Fatal("usage log retry did not persist")
	}

	require.GreaterOrEqual(t, int(repo.calls.Load()), 3)
	last := repo.lastLog.Load()
	require.NotNil(t, last)
	require.Equal(t, "retry-request", last.RequestID)
	require.Nil(t, last.User)
	require.Nil(t, last.APIKey)
	require.Nil(t, last.Account)
	require.NotSame(t, usageLog.UpstreamModel, last.UpstreamModel)
	require.Equal(t, upstreamModel, *last.UpstreamModel)
}
