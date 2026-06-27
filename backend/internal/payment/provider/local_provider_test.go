package provider

import (
	"context"
	"strings"
	"testing"

	"github.com/Wei-Shaw/sub2api/internal/payment"
)

func TestNewLocalTestRequiresExplicitEnable(t *testing.T) {
	t.Setenv(enableLocalTestPaymentEnv, "")

	_, err := NewLocalTest("local", nil)
	if err == nil {
		t.Fatal("NewLocalTest should reject when local test payments are not explicitly enabled")
	}
}

func TestLocalTestCreatePaymentWhenEnabled(t *testing.T) {
	t.Setenv(enableLocalTestPaymentEnv, "true")

	prov, err := NewLocalTest("local", nil)
	if err != nil {
		t.Fatalf("NewLocalTest returned error: %v", err)
	}

	resp, err := prov.CreatePayment(context.Background(), payment.CreatePaymentRequest{
		OrderID: "sub2_test",
		Amount:  "59.99",
	})
	if err != nil {
		t.Fatalf("CreatePayment returned error: %v", err)
	}
	if resp.TradeNo != "sub2_test" {
		t.Fatalf("TradeNo = %q, want sub2_test", resp.TradeNo)
	}
	if !strings.HasPrefix(resp.QRCode, "weixin://wxpay/local-test?") {
		t.Fatalf("QRCode = %q, want local test weixin URL", resp.QRCode)
	}
}
