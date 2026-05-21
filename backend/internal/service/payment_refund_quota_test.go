package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProratedRefundQuotaDefaultSuite(t *testing.T) {
	tests := []struct {
		name         string
		totalQuota   int
		orderAmount  float64
		refundAmount float64
		want         int
	}{
		{name: "full refund returns full quota", totalQuota: 1000, orderAmount: 29.9, refundAmount: 29.9, want: 1000},
		{name: "partial refund rounds up", totalQuota: 1000, orderAmount: 100, refundAmount: 33.1, want: 331},
		{name: "tiny partial refund deducts at least one call", totalQuota: 1000, orderAmount: 100, refundAmount: 0.01, want: 1},
		{name: "over refund caps at full quota", totalQuota: 1000, orderAmount: 100, refundAmount: 120, want: 1000},
		{name: "zero quota returns zero", totalQuota: 0, orderAmount: 100, refundAmount: 50, want: 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.want, proratedRefundQuota(tt.totalQuota, tt.orderAmount, tt.refundAmount))
		})
	}
}
