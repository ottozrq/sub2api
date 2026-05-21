package repository

import (
	"context"
	"errors"
	"testing"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	"github.com/Wei-Shaw/sub2api/internal/service"
	"github.com/stretchr/testify/require"
)

func TestEnqueueUsageLogRetryRequeuesDeadRowsWithFreshAttempts(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherRegexp))
	require.NoError(t, err)
	defer db.Close()

	repo := &usageLogRepository{db: db}
	mock.ExpectExec(`(?s)INSERT INTO usage_log_retry_queue .*ON CONFLICT .*status = 'pending'.*attempts = 0`).
		WithArgs("req-1", int64(42), "service.test", sqlmock.AnyArg(), "insert failed").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.EnqueueUsageLogRetry(context.Background(), &service.UsageLog{
		RequestID: "req-1",
		APIKeyID:  42,
	}, "service.test", errors.New("insert failed"))

	require.NoError(t, err)
	require.NoError(t, mock.ExpectationsWereMet())
}
