package service

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGPT56ModelRoutingAndReasoning(t *testing.T) {
	tests := map[string]string{
		"gpt-5.6":                        "gpt-5.6-sol",
		"gpt-5.6-max":                    "gpt-5.6-sol",
		"gpt-5.6-terra":                  "gpt-5.6-terra",
		"openai/gpt-5.6-terra-high":      "gpt-5.6-terra",
		"gpt-5.6-luna":                   "gpt-5.6-luna",
		"openai/gpt-5.6-luna-2026-07-01": "gpt-5.6-luna",
	}
	for input, want := range tests {
		t.Run(input, func(t *testing.T) {
			require.Equal(t, want, normalizeCodexModel(input))
		})
	}

	require.Equal(t, "max", normalizeOpenAIReasoningEffortForModel("max", "gpt-5.6-terra"))
	require.Equal(t, "max", normalizeOpenAIReasoningEffortForModel("MAX", "openai/gpt-5.6-luna-high"))
	require.Empty(t, normalizeOpenAIReasoningEffortForModel("max", "gpt-5.5"))
}

func TestGPT56OfficialFallbackBilling(t *testing.T) {
	tests := []struct {
		model                                                        string
		input, inputPriority, output, outputPriority                 float64
		cacheWrite, cacheWritePriority, cacheRead, cacheReadPriority float64
	}{
		{"gpt-5.6-terra", 2.5e-6, 5e-6, 15e-6, 30e-6, 3.125e-6, 6.25e-6, 0.25e-6, 0.5e-6},
		{"gpt-5.6-luna", 1e-6, 2e-6, 6e-6, 12e-6, 1.25e-6, 2.5e-6, 0.1e-6, 0.2e-6},
	}

	billing := NewBillingService(nil, nil)
	for _, tt := range tests {
		t.Run(tt.model, func(t *testing.T) {
			pricing, err := billing.GetModelPricing(tt.model)
			require.NoError(t, err)
			require.InDelta(t, tt.input, pricing.InputPricePerToken, 1e-15)
			require.InDelta(t, tt.inputPriority, pricing.InputPricePerTokenPriority, 1e-15)
			require.InDelta(t, tt.output, pricing.OutputPricePerToken, 1e-15)
			require.InDelta(t, tt.outputPriority, pricing.OutputPricePerTokenPriority, 1e-15)
			require.InDelta(t, tt.cacheWrite, pricing.CacheCreationPricePerToken, 1e-15)
			require.InDelta(t, tt.cacheWritePriority, pricing.CacheCreationPricePerTokenPriority, 1e-15)
			require.InDelta(t, tt.cacheRead, pricing.CacheReadPricePerToken, 1e-15)
			require.InDelta(t, tt.cacheReadPriority, pricing.CacheReadPricePerTokenPriority, 1e-15)

			cost, err := billing.CalculateCostWithServiceTier(tt.model, UsageTokens{
				InputTokens: 1, OutputTokens: 1, CacheCreationTokens: 1, CacheReadTokens: 1,
			}, 1, "priority")
			require.NoError(t, err)
			require.InDelta(t, tt.inputPriority, cost.InputCost, 1e-15)
			require.InDelta(t, tt.outputPriority, cost.OutputCost, 1e-15)
			require.InDelta(t, tt.cacheWritePriority, cost.CacheCreationCost, 1e-15)
			require.InDelta(t, tt.cacheReadPriority, cost.CacheReadCost, 1e-15)
		})
	}
}

func TestGPT56ResponsesUsageIncludesCacheWriteTokens(t *testing.T) {
	usage, ok := extractOpenAIUsageFromJSONBytes([]byte(`{
		"usage": {
			"input_tokens": 100,
			"output_tokens": 20,
			"input_tokens_details": {"cached_tokens": 40, "cache_write_tokens": 15}
		}
	}`))
	require.True(t, ok)
	require.Equal(t, 100, usage.InputTokens)
	require.Equal(t, 40, usage.CacheReadInputTokens)
	require.Equal(t, 15, usage.CacheCreationInputTokens)
}
