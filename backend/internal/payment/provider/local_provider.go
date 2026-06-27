package provider

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/Wei-Shaw/sub2api/internal/payment"
)

const enableLocalTestPaymentEnv = "SUB2API_ENABLE_LOCAL_TEST_PAYMENT"

// LocalTest returns deterministic QR-code content for local payment-flow testing.
// It never talks to an upstream gateway and should not be used for real payments.
type LocalTest struct {
	instanceID string
}

func NewLocalTest(instanceID string, _ map[string]string) (*LocalTest, error) {
	if strings.TrimSpace(os.Getenv(enableLocalTestPaymentEnv)) != "true" {
		return nil, fmt.Errorf("local test payment provider is disabled; set %s=true to enable it for local development", enableLocalTestPaymentEnv)
	}
	return &LocalTest{instanceID: instanceID}, nil
}

func (l *LocalTest) Name() string { return "Local Test Pay" }

func (l *LocalTest) ProviderKey() string { return payment.TypeLocalTest }

func (l *LocalTest) SupportedTypes() []payment.PaymentType {
	return []payment.PaymentType{payment.TypeWxpay}
}

func (l *LocalTest) CreatePayment(_ context.Context, req payment.CreatePaymentRequest) (*payment.CreatePaymentResponse, error) {
	values := url.Values{}
	values.Set("order", strings.TrimSpace(req.OrderID))
	values.Set("amount", strings.TrimSpace(req.Amount))
	if l.instanceID != "" {
		values.Set("instance", l.instanceID)
	}
	return &payment.CreatePaymentResponse{
		TradeNo: req.OrderID,
		QRCode:  fmt.Sprintf("weixin://wxpay/local-test?%s", values.Encode()),
	}, nil
}

func (l *LocalTest) QueryOrder(_ context.Context, tradeNo string) (*payment.QueryOrderResponse, error) {
	return &payment.QueryOrderResponse{
		TradeNo: tradeNo,
		Status:  payment.ProviderStatusPending,
	}, nil
}

func (l *LocalTest) VerifyNotification(_ context.Context, _ string, _ map[string]string) (*payment.PaymentNotification, error) {
	return nil, nil
}

func (l *LocalTest) Refund(_ context.Context, req payment.RefundRequest) (*payment.RefundResponse, error) {
	return &payment.RefundResponse{
		RefundID: "local_test_" + strings.TrimSpace(req.OrderID),
		Status:   payment.ProviderStatusPending,
	}, nil
}
