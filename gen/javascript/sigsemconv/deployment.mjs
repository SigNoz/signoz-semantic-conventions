//----------------------------------------------------------------------------
// Code generated from semantic convention specification. DO NOT EDIT.
//----------------------------------------------------------------------------

/**
 * A deployment of a service.
 *
 * @note `service.name`, `service.version` and `deployment.environment.name` are set
 * on the OpenTelemetry Resource; every other attribute is set on the log
 * record. Severity is `INFO` (9) for `succeeded` and `ERROR` (17) for `failed`.
 *
 * @experimental This event is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
export const EVENT_SIGNOZ_DEPLOYMENT = 'signoz.deployment';
