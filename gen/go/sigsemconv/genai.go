// Code generated from semantic convention specification. DO NOT EDIT.

package sigsemconv

import "go.opentelemetry.io/otel/attribute"

// Area: genai
const (
	// SigNozGenAIUsageCacheReadInputTokensCostKey is the attribute Key conforming
	// to the "signoz.gen_ai.usage.cache_read.input_tokens.cost" semantic
	// conventions. It represents the monetary cost of the input tokens served from
	// a provider-managed cache.
	//
	// Type: double
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: 3e-5
	// Note: `gen_ai.usage.cache_read.input_tokens` with `.cost` appended. The value
	// SHOULD be included in `signoz.gen_ai.usage.tokens.cost`.
	//
	// Providers price cache reads below their standard input rate.
	SigNozGenAIUsageCacheReadInputTokensCostKey = attribute.Key("signoz.gen_ai.usage.cache_read.input_tokens.cost")

	// SigNozGenAIUsageCacheWriteInputTokensCostKey is the attribute Key conforming
	// to the "signoz.gen_ai.usage.cache_write.input_tokens.cost" semantic
	// conventions. It represents the monetary cost of the input tokens written to a
	// provider-managed cache.
	//
	// Type: double
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: 0.00094
	// Note: `gen_ai.usage.cache_write.input_tokens` with `.cost` appended. The
	// value
	// SHOULD be included in `signoz.gen_ai.usage.tokens.cost`.
	//
	// Providers price cache writes above their standard input rate.
	SigNozGenAIUsageCacheWriteInputTokensCostKey = attribute.Key("signoz.gen_ai.usage.cache_write.input_tokens.cost")

	// SigNozGenAIUsageInputTokensCostKey is the attribute Key conforming to the
	// "signoz.gen_ai.usage.input_tokens.cost" semantic conventions. It represents
	// the monetary cost of the input tokens billed at the model's standard input
	// rate.
	//
	// Type: double
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: 0.00075
	// Note: `gen_ai.usage.input_tokens` with `.cost` appended. The value SHOULD be
	// included in `signoz.gen_ai.usage.tokens.cost`.
	//
	// Tokens read from or written to a provider-managed cache are billed at
	// their own rates and are excluded here, even though
	// `gen_ai.usage.input_tokens` counts them.
	SigNozGenAIUsageInputTokensCostKey = attribute.Key("signoz.gen_ai.usage.input_tokens.cost")

	// SigNozGenAIUsageOutputTokensCostKey is the attribute Key conforming to the
	// "signoz.gen_ai.usage.output_tokens.cost" semantic conventions. It represents
	// the monetary cost of the output tokens a model generated.
	//
	// Type: double
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: 0.00159
	// Note: `gen_ai.usage.output_tokens` with `.cost` appended. The value SHOULD be
	// included in `signoz.gen_ai.usage.tokens.cost`.
	//
	// Reasoning tokens are part of the output and are priced with it.
	SigNozGenAIUsageOutputTokensCostKey = attribute.Key("signoz.gen_ai.usage.output_tokens.cost")

	// SigNozGenAIUsageTokensCostKey is the attribute Key conforming to the
	// "signoz.gen_ai.usage.tokens.cost" semantic conventions. It represents the
	// total monetary cost of the tokens a model call used.
	//
	// Type: double
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: 0.00234, 0.4812
	// Note: The sum of the per-class costs. Upstream has no count for all classes
	// together, so this name has no `gen_ai.usage.*` counterpart.
	//
	// This value SHOULD include all charges for the tokens on the call: the
	// input and output tokens, and any tokens read from or written to a cache.
	//
	// The per-class costs are disjoint and sum to this value. They do not nest
	// the way the counts they are derived from do, where
	// `gen_ai.usage.cache_read.input_tokens` is a subset of
	// `gen_ai.usage.input_tokens`; tokens are priced in exactly one class.
	SigNozGenAIUsageTokensCostKey = attribute.Key("signoz.gen_ai.usage.tokens.cost")
)

// SigNozGenAIUsageCacheReadInputTokensCost returns an attribute KeyValue
// conforming to the "signoz.gen_ai.usage.cache_read.input_tokens.cost" semantic
// conventions. It represents the monetary cost of the input tokens served from a
// provider-managed cache.
func SigNozGenAIUsageCacheReadInputTokensCost(val float64) attribute.KeyValue {
	return SigNozGenAIUsageCacheReadInputTokensCostKey.Float64(val)
}

// SigNozGenAIUsageCacheWriteInputTokensCost returns an attribute KeyValue
// conforming to the "signoz.gen_ai.usage.cache_write.input_tokens.cost" semantic
// conventions. It represents the monetary cost of the input tokens written to a
// provider-managed cache.
func SigNozGenAIUsageCacheWriteInputTokensCost(val float64) attribute.KeyValue {
	return SigNozGenAIUsageCacheWriteInputTokensCostKey.Float64(val)
}

// SigNozGenAIUsageInputTokensCost returns an attribute KeyValue conforming to
// the "signoz.gen_ai.usage.input_tokens.cost" semantic conventions. It
// represents the monetary cost of the input tokens billed at the model's
// standard input rate.
func SigNozGenAIUsageInputTokensCost(val float64) attribute.KeyValue {
	return SigNozGenAIUsageInputTokensCostKey.Float64(val)
}

// SigNozGenAIUsageOutputTokensCost returns an attribute KeyValue conforming to
// the "signoz.gen_ai.usage.output_tokens.cost" semantic conventions. It
// represents the monetary cost of the output tokens a model generated.
func SigNozGenAIUsageOutputTokensCost(val float64) attribute.KeyValue {
	return SigNozGenAIUsageOutputTokensCostKey.Float64(val)
}

// SigNozGenAIUsageTokensCost returns an attribute KeyValue conforming to the
// "signoz.gen_ai.usage.tokens.cost" semantic conventions. It represents the
// total monetary cost of the tokens a model call used.
func SigNozGenAIUsageTokensCost(val float64) attribute.KeyValue {
	return SigNozGenAIUsageTokensCostKey.Float64(val)
}
