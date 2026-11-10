//----------------------------------------------------------------------------
// Code generated from semantic convention specification. DO NOT EDIT.
//----------------------------------------------------------------------------

/**
 * The monetary cost of the input tokens served from a provider-managed cache.
 *
 * @example 0.00003
 *
 * @note `gen_ai.usage.cache_read.input_tokens` with `.cost` appended. The value
 * **SHOULD** be included in `signoz.gen_ai.usage.tokens.cost`.
 *
 * Providers price cache reads below their standard input rate.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
export const ATTR_SIGNOZ_GEN_AI_USAGE_CACHE_READ_INPUT_TOKENS_COST = 'signoz.gen_ai.usage.cache_read.input_tokens.cost';

/**
 * The monetary cost of the input tokens written to a provider-managed cache.
 *
 * @example 0.00094
 *
 * @note `gen_ai.usage.cache_write.input_tokens` with `.cost` appended. The value
 * **SHOULD** be included in `signoz.gen_ai.usage.tokens.cost`.
 *
 * Providers price cache writes above their standard input rate.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
export const ATTR_SIGNOZ_GEN_AI_USAGE_CACHE_WRITE_INPUT_TOKENS_COST = 'signoz.gen_ai.usage.cache_write.input_tokens.cost';

/**
 * The monetary cost of the input tokens billed at the model's standard input rate.
 *
 * @example 0.00075
 *
 * @note `gen_ai.usage.input_tokens` with `.cost` appended. The value **SHOULD** be
 * included in `signoz.gen_ai.usage.tokens.cost`.
 *
 * Tokens read from or written to a provider-managed cache are billed at
 * their own rates and are excluded here, even though
 * `gen_ai.usage.input_tokens` counts them.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
export const ATTR_SIGNOZ_GEN_AI_USAGE_INPUT_TOKENS_COST = 'signoz.gen_ai.usage.input_tokens.cost';

/**
 * The monetary cost of the output tokens a model generated.
 *
 * @example 0.00159
 *
 * @note `gen_ai.usage.output_tokens` with `.cost` appended. The value **SHOULD** be
 * included in `signoz.gen_ai.usage.tokens.cost`.
 *
 * Reasoning tokens are part of the output and are priced with it.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
export const ATTR_SIGNOZ_GEN_AI_USAGE_OUTPUT_TOKENS_COST = 'signoz.gen_ai.usage.output_tokens.cost';

/**
 * The total monetary cost of the tokens a model call used.
 *
 * @example 0.00234
 * @example 0.4812
 *
 * @note The sum of the per-class costs. Upstream has no count for all classes
 * together, so this name has no `gen_ai.usage.*` counterpart.
 *
 * This value **SHOULD** include all charges for the tokens on the call: the
 * input and output tokens, and any tokens read from or written to a cache.
 *
 * The per-class costs are disjoint and sum to this value. They do not nest
 * the way the counts they are derived from do, where
 * `gen_ai.usage.cache_read.input_tokens` is a subset of
 * `gen_ai.usage.input_tokens`; tokens are priced in exactly one class.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
export const ATTR_SIGNOZ_GEN_AI_USAGE_TOKENS_COST = 'signoz.gen_ai.usage.tokens.cost';
