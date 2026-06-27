package service

import "testing"

func TestValidateOrderInputBalanceRequiresCreditAmount(t *testing.T) {
	svc := &PaymentService{}
	_, err := svc.validateOrderInput(t.Context(), CreateOrderRequest{
		Amount:    10,
		OrderType: "balance",
	}, &PaymentConfig{})
	if err == nil {
		t.Fatal("validateOrderInput should reject balance recharge without credit amount")
	}
}

func TestValidateOrderInputRejectsInvalidOrderType(t *testing.T) {
	svc := &PaymentService{}
	_, err := svc.validateOrderInput(t.Context(), CreateOrderRequest{
		Amount:              10,
		OrderType:           "topup",
		BalanceCreditAmount: 10,
	}, &PaymentConfig{})
	if err == nil {
		t.Fatal("validateOrderInput should reject unsupported order types")
	}
}

func TestCalculateBalanceRechargePaymentAmount(t *testing.T) {
	tests := []struct {
		credit float64
		want   float64
	}{
		{credit: 10, want: 59.99},
		{credit: 20, want: 119.99},
		{credit: 50, want: 299.99},
	}

	for _, tt := range tests {
		if got := calculateBalanceRechargePaymentAmount(tt.credit); got != tt.want {
			t.Fatalf("calculateBalanceRechargePaymentAmount(%v) = %v, want %v", tt.credit, got, tt.want)
		}
	}
}

func TestBalanceRechargePaymentMatchesCredit(t *testing.T) {
	if !balanceRechargePaymentMatchesCredit(59.99, 10) {
		t.Fatal("59.99 should match 10 USD credit")
	}
	if balanceRechargePaymentMatchesCredit(60.00, 10) {
		t.Fatal("60.00 should not match 10 USD credit")
	}
}
