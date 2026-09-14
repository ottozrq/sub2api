package service

import (
	"testing"

	openaiapi "github.com/Wei-Shaw/sub2api/internal/pkg/openai"
	"github.com/stretchr/testify/require"
)

func TestGPT6AstraCatalogRoutingAndReasoning(t *testing.T) {
	ids := openaiapi.DefaultModelIDs()
	require.Contains(t, ids, "gpt-6")
	require.Contains(t, ids, "gpt-6-astra")

	for input, want := range map[string]string{
		"gpt-6":                         "gpt-6-astra",
		"gpt-6-astra":                   "gpt-6-astra",
		"openai/gpt-6-astra-2026-09-03": "gpt-6-astra",
	} {
		t.Run(input, func(t *testing.T) {
			require.Equal(t, want, normalizeCodexModel(input))
		})
	}

	require.Equal(t, "xhigh", normalizeOpenAIReasoningEffortForModel("xhigh", "gpt-6-astra"))
	require.Equal(t, "max", normalizeOpenAIReasoningEffortForModel("max", "gpt-6-astra"))
	require.Empty(t, normalizeOpenAIReasoningEffortForModel("none", "gpt-6-astra"))
}

func TestGPT6AstraOfficialFallbackBilling(t *testing.T) {
	billing := NewBillingService(nil, nil)
	for _, model := range []string{"gpt-6", "gpt-6-astra", "openai/gpt-6-astra-2026-09-03"} {
		t.Run(model, func(t *testing.T) {
			pricing, err := billing.GetModelPricing(model)
			require.NoError(t, err)
			require.InDelta(t, 10e-6, pricing.InputPricePerToken, 1e-15)
			require.InDelta(t, 20e-6, pricing.InputPricePerTokenPriority, 1e-15)
			require.InDelta(t, 50e-6, pricing.OutputPricePerToken, 1e-15)
			require.InDelta(t, 100e-6, pricing.OutputPricePerTokenPriority, 1e-15)
			require.InDelta(t, 12.5e-6, pricing.CacheCreationPricePerToken, 1e-15)
			require.InDelta(t, 25e-6, pricing.CacheCreationPricePerTokenPriority, 1e-15)
			require.InDelta(t, 1e-6, pricing.CacheReadPricePerToken, 1e-15)
			require.InDelta(t, 2e-6, pricing.CacheReadPricePerTokenPriority, 1e-15)
			require.Equal(t, 272_000, pricing.LongContextInputThreshold)
			require.InDelta(t, 2.0, pricing.LongContextInputMultiplier, 1e-15)
			require.InDelta(t, 1.5, pricing.LongContextOutputMultiplier, 1e-15)
		})
	}
}

func TestGPT6AstraPricingServiceFallback(t *testing.T) {
	require.Equal(t, "gpt-6-astra", normalizeModelNameForPricing("gpt-6"))
	require.Equal(t, "gpt-6-astra-2026-09-03", normalizeModelNameForPricing("openai/gpt-6-astra-2026-09-03"))

	for _, model := range []string{"gpt-6", "gpt-6-astra", "openai/gpt-6-astra-2026-09-03"} {
		pricing := (&PricingService{}).matchOpenAIModel(normalizeModelNameForPricing(model))
		require.NotNil(t, pricing)
		require.InDelta(t, 10e-6, pricing.InputCostPerToken, 1e-15)
		require.InDelta(t, 50e-6, pricing.OutputCostPerToken, 1e-15)
	}
}
