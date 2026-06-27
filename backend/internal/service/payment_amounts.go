package service

import (
	"math"

	"github.com/shopspring/decimal"
)

const defaultBalanceRechargeMultiplier = 1.0
const balanceRechargeCNYPerUSD = 6.0
const balanceRechargeDiscountCNY = 0.01
const balanceRechargeMatchToleranceCNY = 0.005

func normalizeBalanceRechargeMultiplier(multiplier float64) float64 {
	if math.IsNaN(multiplier) || math.IsInf(multiplier, 0) || multiplier <= 0 {
		return defaultBalanceRechargeMultiplier
	}
	return multiplier
}

func calculateCreditedBalance(paymentAmount, multiplier float64) float64 {
	return decimal.NewFromFloat(paymentAmount).
		Mul(decimal.NewFromFloat(normalizeBalanceRechargeMultiplier(multiplier))).
		Round(2).
		InexactFloat64()
}

func calculateBalanceRechargePaymentAmount(creditAmount float64) float64 {
	if creditAmount <= 0 {
		return 0
	}
	amount := decimal.NewFromFloat(creditAmount).
		Mul(decimal.NewFromFloat(balanceRechargeCNYPerUSD)).
		Sub(decimal.NewFromFloat(balanceRechargeDiscountCNY)).
		Round(2).
		InexactFloat64()
	if amount < 0 {
		return 0
	}
	return amount
}

func balanceRechargePaymentMatchesCredit(paymentAmount, creditAmount float64) bool {
	if creditAmount <= 0 {
		return true
	}
	expected := calculateBalanceRechargePaymentAmount(creditAmount)
	return math.Abs(paymentAmount-expected) < balanceRechargeMatchToleranceCNY
}

func calculateGatewayRefundAmount(orderAmount, payAmount, refundAmount float64) float64 {
	if orderAmount <= 0 || payAmount <= 0 || refundAmount <= 0 {
		return 0
	}
	if math.Abs(refundAmount-orderAmount) <= amountToleranceCNY {
		return decimal.NewFromFloat(payAmount).Round(2).InexactFloat64()
	}
	return decimal.NewFromFloat(payAmount).
		Mul(decimal.NewFromFloat(refundAmount)).
		Div(decimal.NewFromFloat(orderAmount)).
		Round(2).
		InexactFloat64()
}
