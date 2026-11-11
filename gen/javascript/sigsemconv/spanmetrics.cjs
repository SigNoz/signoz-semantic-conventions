//----------------------------------------------------------------------------
// Code generated from semantic convention specification. DO NOT EDIT.
//----------------------------------------------------------------------------

'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

/**
 * The remote endpoint of an external call.
 *
 * @example payments.internal:8443
 * @example oteldemo.CartService/GetCart
 *
 * @note `<rpc.service>/<rpc.method>` for RPC calls, otherwise the remote host
 * as `host` or `host:port`, or the span's `peer.service`.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_ADDRESS = 'address';

/**
 * The name of the aggregated spans.
 *
 * @example GET /api/v1/orders/{id}
 * @example overflow_operation
 *
 * @note Set to `overflow_operation` when the service has more distinct span
 * names than are tracked.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_OPERATION = 'operation';

/**
 * The value of `deployment.environment` on the span's Resource.
 *
 * @example production
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_RESOURCE_DEPLOYMENT_ENVIRONMENT = 'resource_deployment.environment';

/**
 * The value of `service.namespace` on the span's Resource.
 *
 * @example shop
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_RESOURCE_SERVICE_NAMESPACE = 'resource_service.namespace';

/**
 * A copy of `signoz.collector.id`.
 *
 * @example 6f1c2b7e-4d3a-4b8f-9e21-5a0c7d9b3e14
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_RESOURCE_SIGNOZ_COLLECTOR_ID = 'resource_signoz.collector.id';

/**
 * The identifier of the collector instance that produced the metric.
 *
 * @example 6f1c2b7e-4d3a-4b8f-9e21-5a0c7d9b3e14
 *
 * @note The collector's `service.instance.id`, or a UUID assigned when the
 * collector starts if it has none.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_COLLECTOR_ID = 'signoz.collector.id';

/**
 * The kind of the aggregated spans.
 *
 * @example SPAN_KIND_SERVER
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SPAN_KIND = 'span.kind';

/**
 * Enum value "SPAN_KIND_CLIENT" for attribute {@link ATTR_SPAN_KIND}.
 *
 * The client side of a remote call.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SPAN_KIND_VALUE_CLIENT = "SPAN_KIND_CLIENT";

/**
 * Enum value "SPAN_KIND_CONSUMER" for attribute {@link ATTR_SPAN_KIND}.
 *
 * The consumer of an asynchronous message.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SPAN_KIND_VALUE_CONSUMER = "SPAN_KIND_CONSUMER";

/**
 * Enum value "SPAN_KIND_INTERNAL" for attribute {@link ATTR_SPAN_KIND}.
 *
 * An internal operation.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SPAN_KIND_VALUE_INTERNAL = "SPAN_KIND_INTERNAL";

/**
 * Enum value "SPAN_KIND_PRODUCER" for attribute {@link ATTR_SPAN_KIND}.
 *
 * The producer of an asynchronous message.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SPAN_KIND_VALUE_PRODUCER = "SPAN_KIND_PRODUCER";

/**
 * Enum value "SPAN_KIND_SERVER" for attribute {@link ATTR_SPAN_KIND}.
 *
 * The server side of a remote call.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SPAN_KIND_VALUE_SERVER = "SPAN_KIND_SERVER";

/**
 * Enum value "SPAN_KIND_UNSPECIFIED" for attribute {@link ATTR_SPAN_KIND}.
 *
 * The span kind is unspecified.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SPAN_KIND_VALUE_UNSPECIFIED = "SPAN_KIND_UNSPECIFIED";

/**
 * The status code of the aggregated spans.
 *
 * @example STATUS_CODE_ERROR
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_STATUS_CODE = 'status.code';

/**
 * Enum value "STATUS_CODE_ERROR" for attribute {@link ATTR_STATUS_CODE}.
 *
 * The span completed with an error.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.STATUS_CODE_VALUE_ERROR = "STATUS_CODE_ERROR";

/**
 * Enum value "STATUS_CODE_OK" for attribute {@link ATTR_STATUS_CODE}.
 *
 * The span completed successfully.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.STATUS_CODE_VALUE_OK = "STATUS_CODE_OK";

/**
 * Enum value "STATUS_CODE_UNSET" for attribute {@link ATTR_STATUS_CODE}.
 *
 * The span status is unset.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.STATUS_CODE_VALUE_UNSET = "STATUS_CODE_UNSET";

/**
 * The number of spans observed.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_CALLS_TOTAL = 'signoz_calls_total';

/**
 * The number of database calls observed.
 *
 * @note A database call is a span that carries `db.system` and is not of kind
 * server.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_DB_LATENCY_COUNT = 'signoz_db_latency_count';

/**
 * The total duration of database calls observed.
 *
 * @note Values are in milliseconds. A database call is a span that carries
 * `db.system` and is not of kind server.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_DB_LATENCY_SUM = 'signoz_db_latency_sum';

/**
 * The number of external calls observed.
 *
 * @note An external call is a client span with a known remote endpoint.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_EXTERNAL_CALL_LATENCY_COUNT = 'signoz_external_call_latency_count';

/**
 * The total duration of external calls observed.
 *
 * @note Values are in milliseconds. An external call is a client span with a
 * known remote endpoint.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_EXTERNAL_CALL_LATENCY_SUM = 'signoz_external_call_latency_sum';

/**
 * The duration of spans observed.
 *
 * @note Recorded as an explicit bucket histogram and, when enabled, also as an
 * exponential histogram under the same name. Values are in milliseconds.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_LATENCY = 'signoz_latency';
