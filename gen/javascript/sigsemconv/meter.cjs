//----------------------------------------------------------------------------
// Code generated from semantic convention specification. DO NOT EDIT.
//----------------------------------------------------------------------------

'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

/**
 * The identifier of the meter connector instance that produced the metric.
 *
 * @example 0b8f4a6e-2c1d-4e5f-9a7b-3c6d8e1f2a4b
 *
 * @note A UUID assigned when the connector instance is created.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_CONNECTOR_ID = 'connector_id';

/**
 * The SigNoz workload a Kubernetes node is dedicated to.
 *
 * @example clickhouse
 *
 * @note The value of the node's `signoz.io/workload` label.
 *
 * Known values: `clickhouse`, `redpanda`, `consumers`, `zookeeper`,
 * `gateway`, `signoz`.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_IO_WORKLOAD = 'signoz.io/workload';

/**
 * The amount of CPU allocatable on a Kubernetes node that runs SigNoz.
 *
 * @note The value is in CPU cores and is recorded once per node every hour.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_METER_CPU_NODE_ALLOCATABLE_USAGE = 'signoz.meter.cpu.node.allocatable.usage';

/**
 * The number of log records observed.
 *
 * @note Each log record counts as one.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_METER_LOG_COUNT = 'signoz.meter.log.count';

/**
 * The size of log records observed.
 *
 * @note The size of a log record is the size of its resource attributes, its
 * attributes and its body. `signoz.workspace.*` attributes are not
 * counted.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_METER_LOG_SIZE = 'signoz.meter.log.size';

/**
 * The number of metric data points indexed.
 *
 * @note A gauge or sum data point counts as one. A histogram data point counts
 * as its number of buckets plus one for its count and one each for its
 * sum, min and max when present.
 * A summary data point counts as its number of quantiles plus two.
 * Exponential histogram data points, and metrics whose name starts with
 * `signoz` or `otelcol`, are not counted.
 * Data points that a metric reduction rule condenses are not counted.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_METER_METRIC_DATAPOINT_COUNT = 'signoz.meter.metric.datapoint.count';

/**
 * The number of metric data points ingested.
 *
 * @note A gauge or sum data point counts as one. A histogram data point counts
 * as its number of buckets plus one for its count and one each for its
 * sum, min and max when present.
 * A summary data point counts as its number of quantiles plus two.
 * Exponential histogram data points, and metrics whose name starts with
 * `signoz` or `otelcol`, are not counted.
 * Data points that a metric reduction rule condenses are counted.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_METER_METRIC_DATAPOINT_INGESTED_COUNT = 'signoz.meter.metric.datapoint.ingested.count';

/**
 * The size of metric data points observed.
 *
 * @note The value is always `0`.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_METER_METRIC_DATAPOINT_SIZE = 'signoz.meter.metric.datapoint.size';

/**
 * The number of spans observed.
 *
 * @note Each span counts as one.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_METER_SPAN_COUNT = 'signoz.meter.span.count';

/**
 * The size of spans observed.
 *
 * @note The size of a span is the size of its fields, its resource attributes,
 * its attributes, its events and its links. `signoz.workspace.*` and
 * `signoz.gen_ai.*` attributes are not counted.
 *
 * @experimental This metric is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.METRIC_SIGNOZ_METER_SPAN_SIZE = 'signoz.meter.span.size';
