package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreatePlanRejectsStandardGroup(t *testing.T) {
	ctx := context.Background()
	client := newPaymentConfigServiceTestClient(t)
	svc := &PaymentConfigService{entClient: client}

	standardGroup, err := client.Group.Create().
		SetName("standard").
		SetSubscriptionType(SubscriptionTypeStandard).
		Save(ctx)
	require.NoError(t, err)

	_, err = svc.CreatePlan(ctx, CreatePlanRequest{
		GroupID:      standardGroup.ID,
		Name:         "Pro",
		Price:        9.99,
		ValidityDays: 30,
		ValidityUnit: "days",
		ForSale:      true,
	})
	require.Error(t, err)
	require.Contains(t, err.Error(), "plan group must be a subscription type")
}

func TestUpdatePlanRejectsSwitchingToStandardGroup(t *testing.T) {
	ctx := context.Background()
	client := newPaymentConfigServiceTestClient(t)
	svc := &PaymentConfigService{entClient: client}

	subGroup, err := client.Group.Create().
		SetName("subscription").
		SetSubscriptionType(SubscriptionTypeSubscription).
		Save(ctx)
	require.NoError(t, err)
	standardGroup, err := client.Group.Create().
		SetName("standard").
		SetSubscriptionType(SubscriptionTypeStandard).
		Save(ctx)
	require.NoError(t, err)

	plan, err := svc.CreatePlan(ctx, CreatePlanRequest{
		GroupID:      subGroup.ID,
		Name:         "Pro",
		Price:        9.99,
		ValidityDays: 30,
		ValidityUnit: "days",
		ForSale:      true,
	})
	require.NoError(t, err)

	_, err = svc.UpdatePlan(ctx, plan.ID, UpdatePlanRequest{GroupID: &standardGroup.ID})
	require.Error(t, err)
	require.Contains(t, err.Error(), "plan group must be a subscription type")
}
