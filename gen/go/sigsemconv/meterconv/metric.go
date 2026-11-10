// Code generated from semantic convention specification. DO NOT EDIT.

// Package meterconv provides types and functionality for the SigNoz semantic
// conventions in the "meter" area.
package meterconv

import (
	"context"

	"github.com/SigNoz/signoz-semantic-conventions/gen/go/sigsemconv/internal/metricpool"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

// CPUNodeAllocatableUsage is an instrument used to record metric values
// conforming to the "signoz.meter.cpu.node.allocatable.usage" semantic
// conventions. It represents the amount of CPU allocatable on a Kubernetes node
// that runs SigNoz.
type CPUNodeAllocatableUsage struct {
	metric.Float64Gauge
}

var newCPUNodeAllocatableUsageOpts = []metric.Float64GaugeOption{
	metric.WithDescription("The amount of CPU allocatable on a Kubernetes node that runs SigNoz."),
	metric.WithUnit("{cpu}"),
}

// NewCPUNodeAllocatableUsage returns a new CPUNodeAllocatableUsage instrument.
func NewCPUNodeAllocatableUsage(
	m metric.Meter,
	opt ...metric.Float64GaugeOption,
) (CPUNodeAllocatableUsage, error) {
	// Check if the meter is nil.
	if m == nil {
		return CPUNodeAllocatableUsage{noop.Float64Gauge{}}, nil
	}

	if len(opt) == 0 {
		opt = newCPUNodeAllocatableUsageOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newCPUNodeAllocatableUsageOpts...)
	}

	i, err := m.Float64Gauge(
		"signoz.meter.cpu.node.allocatable.usage",
		opt...,
	)
	if err != nil {
		return CPUNodeAllocatableUsage{noop.Float64Gauge{}}, err
	}
	return CPUNodeAllocatableUsage{i}, nil
}

// Inst returns the underlying metric instrument.
func (m CPUNodeAllocatableUsage) Inst() metric.Float64Gauge {
	return m.Float64Gauge
}

// Name returns the semantic convention name of the instrument.
func (CPUNodeAllocatableUsage) Name() string {
	return "signoz.meter.cpu.node.allocatable.usage"
}

// Unit returns the semantic convention unit of the instrument
func (CPUNodeAllocatableUsage) Unit() string {
	return "{cpu}"
}

// Description returns the semantic convention description of the instrument
func (CPUNodeAllocatableUsage) Description() string {
	return "The amount of CPU allocatable on a Kubernetes node that runs SigNoz."
}

// Record records val to the current distribution for attrs.
//
// The k8sNodeName is the name of the Node.
//
// The k8sNodeUid is the UID of the Node.
//
// The signozIoWorkload is the SigNoz workload a Kubernetes node is dedicated to.
//
// The value is in CPU cores and is recorded once per node every hour.
func (m CPUNodeAllocatableUsage) Record(
	ctx context.Context,
	val float64,
	k8sNodeName string,
	k8sNodeUid string,
	signozIoWorkload string,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Gauge.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Gauge.Record(ctx, val, metric.WithAttributes(
			attribute.String("k8s.node.name", k8sNodeName),
			attribute.String("k8s.node.uid", k8sNodeUid),
			attribute.String("signoz.io/workload", signozIoWorkload),
		))
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("k8s.node.name", k8sNodeName),
				attribute.String("k8s.node.uid", k8sNodeUid),
				attribute.String("signoz.io/workload", signozIoWorkload),
			)...,
		),
	)

	m.Float64Gauge.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// The value is in CPU cores and is recorded once per node every hour.
func (m CPUNodeAllocatableUsage) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Gauge.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Gauge.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Gauge.Record(ctx, val, *o...)
}

// CPUNodeAllocatableUsageObservable is an instrument used to record metric
// values conforming to the "signoz.meter.cpu.node.allocatable.usage" semantic
// conventions. It represents the amount of CPU allocatable on a Kubernetes node
// that runs SigNoz.
type CPUNodeAllocatableUsageObservable struct {
	metric.Float64ObservableGauge
}

var newCPUNodeAllocatableUsageObservableOpts = []metric.Float64ObservableGaugeOption{
	metric.WithDescription("The amount of CPU allocatable on a Kubernetes node that runs SigNoz."),
	metric.WithUnit("{cpu}"),
}

// NewCPUNodeAllocatableUsageObservable returns a new
// CPUNodeAllocatableUsageObservable instrument.
func NewCPUNodeAllocatableUsageObservable(
	m metric.Meter,
	opt ...metric.Float64ObservableGaugeOption,
) (CPUNodeAllocatableUsageObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return CPUNodeAllocatableUsageObservable{noop.Float64ObservableGauge{}}, nil
	}

	if len(opt) == 0 {
		opt = newCPUNodeAllocatableUsageObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newCPUNodeAllocatableUsageObservableOpts...)
	}

	i, err := m.Float64ObservableGauge(
		"signoz.meter.cpu.node.allocatable.usage",
		opt...,
	)
	if err != nil {
		return CPUNodeAllocatableUsageObservable{noop.Float64ObservableGauge{}}, err
	}
	return CPUNodeAllocatableUsageObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m CPUNodeAllocatableUsageObservable) Inst() metric.Float64ObservableGauge {
	return m.Float64ObservableGauge
}

// Name returns the semantic convention name of the instrument.
func (CPUNodeAllocatableUsageObservable) Name() string {
	return "signoz.meter.cpu.node.allocatable.usage"
}

// Unit returns the semantic convention unit of the instrument
func (CPUNodeAllocatableUsageObservable) Unit() string {
	return "{cpu}"
}

// Description returns the semantic convention description of the instrument
func (CPUNodeAllocatableUsageObservable) Description() string {
	return "The amount of CPU allocatable on a Kubernetes node that runs SigNoz."
}

// AttrK8SNodeName returns a required attribute for the "k8s.node.name" semantic
// convention. It represents the name of the Node.
func (CPUNodeAllocatableUsageObservable) AttrK8SNodeName(val string) attribute.KeyValue {
	return attribute.String("k8s.node.name", val)
}

// AttrK8SNodeUID returns a required attribute for the "k8s.node.uid" semantic
// convention. It represents the UID of the Node.
func (CPUNodeAllocatableUsageObservable) AttrK8SNodeUID(val string) attribute.KeyValue {
	return attribute.String("k8s.node.uid", val)
}

// AttrSigNozIOWorkload returns a required attribute for the "signoz.io/workload"
// semantic convention. It represents the SigNoz workload a Kubernetes node is
// dedicated to.
func (CPUNodeAllocatableUsageObservable) AttrSigNozIOWorkload(val string) attribute.KeyValue {
	return attribute.String("signoz.io/workload", val)
}

// LogCount is an instrument used to record metric values conforming to the
// "signoz.meter.log.count" semantic conventions. It represents the number of log
// records observed.
type LogCount struct {
	metric.Int64Counter
}

var newLogCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of log records observed."),
	metric.WithUnit("1"),
}

// NewLogCount returns a new LogCount instrument.
func NewLogCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (LogCount, error) {
	// Check if the meter is nil.
	if m == nil {
		return LogCount{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newLogCountOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newLogCountOpts...)
	}

	i, err := m.Int64Counter(
		"signoz.meter.log.count",
		opt...,
	)
	if err != nil {
		return LogCount{noop.Int64Counter{}}, err
	}
	return LogCount{i}, nil
}

// Inst returns the underlying metric instrument.
func (m LogCount) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (LogCount) Name() string {
	return "signoz.meter.log.count"
}

// Unit returns the semantic convention unit of the instrument
func (LogCount) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (LogCount) Description() string {
	return "The number of log records observed."
}

// Add adds incr to the existing count for attrs.
//
// The connectorId is the identifier of the meter connector instance that
// produced the metric.
//
// All additional attrs passed are included in the recorded value.
//
// Each log record counts as one.
func (m LogCount) Add(
	ctx context.Context,
	incr int64,
	connectorId string,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("connector_id", connectorId),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("connector_id", connectorId),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// Each log record counts as one.
func (m LogCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (LogCount) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (LogCount) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (LogCount) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// LogCountObservable is an instrument used to record metric values conforming to
// the "signoz.meter.log.count" semantic conventions. It represents the number of
// log records observed.
type LogCountObservable struct {
	metric.Int64ObservableCounter
}

var newLogCountObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of log records observed."),
	metric.WithUnit("1"),
}

// NewLogCountObservable returns a new LogCountObservable instrument.
func NewLogCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (LogCountObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return LogCountObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newLogCountObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newLogCountObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz.meter.log.count",
		opt...,
	)
	if err != nil {
		return LogCountObservable{noop.Int64ObservableCounter{}}, err
	}
	return LogCountObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m LogCountObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (LogCountObservable) Name() string {
	return "signoz.meter.log.count"
}

// Unit returns the semantic convention unit of the instrument
func (LogCountObservable) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (LogCountObservable) Description() string {
	return "The number of log records observed."
}

// AttrConnectorID returns a required attribute for the "connector_id" semantic
// convention. It represents the identifier of the meter connector instance that
// produced the metric.
func (LogCountObservable) AttrConnectorID(val string) attribute.KeyValue {
	return attribute.String("connector_id", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (LogCountObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (LogCountObservable) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (LogCountObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// LogSize is an instrument used to record metric values conforming to the
// "signoz.meter.log.size" semantic conventions. It represents the size of log
// records observed.
type LogSize struct {
	metric.Int64Counter
}

var newLogSizeOpts = []metric.Int64CounterOption{
	metric.WithDescription("The size of log records observed."),
	metric.WithUnit("By"),
}

// NewLogSize returns a new LogSize instrument.
func NewLogSize(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (LogSize, error) {
	// Check if the meter is nil.
	if m == nil {
		return LogSize{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newLogSizeOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newLogSizeOpts...)
	}

	i, err := m.Int64Counter(
		"signoz.meter.log.size",
		opt...,
	)
	if err != nil {
		return LogSize{noop.Int64Counter{}}, err
	}
	return LogSize{i}, nil
}

// Inst returns the underlying metric instrument.
func (m LogSize) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (LogSize) Name() string {
	return "signoz.meter.log.size"
}

// Unit returns the semantic convention unit of the instrument
func (LogSize) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (LogSize) Description() string {
	return "The size of log records observed."
}

// Add adds incr to the existing count for attrs.
//
// The connectorId is the identifier of the meter connector instance that
// produced the metric.
//
// All additional attrs passed are included in the recorded value.
//
// The size of a log record is the size of its resource attributes, its
// attributes and its body. `signoz.workspace.*` attributes are not
// counted.
func (m LogSize) Add(
	ctx context.Context,
	incr int64,
	connectorId string,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("connector_id", connectorId),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("connector_id", connectorId),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// The size of a log record is the size of its resource attributes, its
// attributes and its body. `signoz.workspace.*` attributes are not
// counted.
func (m LogSize) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (LogSize) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (LogSize) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (LogSize) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// LogSizeObservable is an instrument used to record metric values conforming to
// the "signoz.meter.log.size" semantic conventions. It represents the size of
// log records observed.
type LogSizeObservable struct {
	metric.Int64ObservableCounter
}

var newLogSizeObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The size of log records observed."),
	metric.WithUnit("By"),
}

// NewLogSizeObservable returns a new LogSizeObservable instrument.
func NewLogSizeObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (LogSizeObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return LogSizeObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newLogSizeObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newLogSizeObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz.meter.log.size",
		opt...,
	)
	if err != nil {
		return LogSizeObservable{noop.Int64ObservableCounter{}}, err
	}
	return LogSizeObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m LogSizeObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (LogSizeObservable) Name() string {
	return "signoz.meter.log.size"
}

// Unit returns the semantic convention unit of the instrument
func (LogSizeObservable) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (LogSizeObservable) Description() string {
	return "The size of log records observed."
}

// AttrConnectorID returns a required attribute for the "connector_id" semantic
// convention. It represents the identifier of the meter connector instance that
// produced the metric.
func (LogSizeObservable) AttrConnectorID(val string) attribute.KeyValue {
	return attribute.String("connector_id", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (LogSizeObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (LogSizeObservable) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (LogSizeObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// MetricDatapointCount is an instrument used to record metric values conforming
// to the "signoz.meter.metric.datapoint.count" semantic conventions. It
// represents the number of metric data points observed.
type MetricDatapointCount struct {
	metric.Int64Counter
}

var newMetricDatapointCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of metric data points observed."),
	metric.WithUnit("1"),
}

// NewMetricDatapointCount returns a new MetricDatapointCount instrument.
func NewMetricDatapointCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (MetricDatapointCount, error) {
	// Check if the meter is nil.
	if m == nil {
		return MetricDatapointCount{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newMetricDatapointCountOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newMetricDatapointCountOpts...)
	}

	i, err := m.Int64Counter(
		"signoz.meter.metric.datapoint.count",
		opt...,
	)
	if err != nil {
		return MetricDatapointCount{noop.Int64Counter{}}, err
	}
	return MetricDatapointCount{i}, nil
}

// Inst returns the underlying metric instrument.
func (m MetricDatapointCount) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (MetricDatapointCount) Name() string {
	return "signoz.meter.metric.datapoint.count"
}

// Unit returns the semantic convention unit of the instrument
func (MetricDatapointCount) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (MetricDatapointCount) Description() string {
	return "The number of metric data points observed."
}

// Add adds incr to the existing count for attrs.
//
// The connectorId is the identifier of the meter connector instance that
// produced the metric.
//
// All additional attrs passed are included in the recorded value.
//
// A gauge or sum data point counts as one. A histogram data point counts
// as its number of buckets plus one for its count and one each for its
// sum, min and max when present.
// A summary data point counts as its number of quantiles plus two.
// Exponential histogram data points, and metrics whose name starts with
// `signoz` or `otelcol`, are not counted.
func (m MetricDatapointCount) Add(
	ctx context.Context,
	incr int64,
	connectorId string,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("connector_id", connectorId),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("connector_id", connectorId),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// A gauge or sum data point counts as one. A histogram data point counts
// as its number of buckets plus one for its count and one each for its
// sum, min and max when present.
// A summary data point counts as its number of quantiles plus two.
// Exponential histogram data points, and metrics whose name starts with
// `signoz` or `otelcol`, are not counted.
func (m MetricDatapointCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (MetricDatapointCount) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (MetricDatapointCount) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (MetricDatapointCount) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// MetricDatapointCountObservable is an instrument used to record metric values
// conforming to the "signoz.meter.metric.datapoint.count" semantic conventions.
// It represents the number of metric data points observed.
type MetricDatapointCountObservable struct {
	metric.Int64ObservableCounter
}

var newMetricDatapointCountObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of metric data points observed."),
	metric.WithUnit("1"),
}

// NewMetricDatapointCountObservable returns a new MetricDatapointCountObservable
// instrument.
func NewMetricDatapointCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (MetricDatapointCountObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return MetricDatapointCountObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newMetricDatapointCountObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newMetricDatapointCountObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz.meter.metric.datapoint.count",
		opt...,
	)
	if err != nil {
		return MetricDatapointCountObservable{noop.Int64ObservableCounter{}}, err
	}
	return MetricDatapointCountObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m MetricDatapointCountObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (MetricDatapointCountObservable) Name() string {
	return "signoz.meter.metric.datapoint.count"
}

// Unit returns the semantic convention unit of the instrument
func (MetricDatapointCountObservable) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (MetricDatapointCountObservable) Description() string {
	return "The number of metric data points observed."
}

// AttrConnectorID returns a required attribute for the "connector_id" semantic
// convention. It represents the identifier of the meter connector instance that
// produced the metric.
func (MetricDatapointCountObservable) AttrConnectorID(val string) attribute.KeyValue {
	return attribute.String("connector_id", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (MetricDatapointCountObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (MetricDatapointCountObservable) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (MetricDatapointCountObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// MetricDatapointSize is an instrument used to record metric values conforming
// to the "signoz.meter.metric.datapoint.size" semantic conventions. It
// represents the size of metric data points observed.
type MetricDatapointSize struct {
	metric.Int64Counter
}

var newMetricDatapointSizeOpts = []metric.Int64CounterOption{
	metric.WithDescription("The size of metric data points observed."),
	metric.WithUnit("By"),
}

// NewMetricDatapointSize returns a new MetricDatapointSize instrument.
func NewMetricDatapointSize(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (MetricDatapointSize, error) {
	// Check if the meter is nil.
	if m == nil {
		return MetricDatapointSize{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newMetricDatapointSizeOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newMetricDatapointSizeOpts...)
	}

	i, err := m.Int64Counter(
		"signoz.meter.metric.datapoint.size",
		opt...,
	)
	if err != nil {
		return MetricDatapointSize{noop.Int64Counter{}}, err
	}
	return MetricDatapointSize{i}, nil
}

// Inst returns the underlying metric instrument.
func (m MetricDatapointSize) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (MetricDatapointSize) Name() string {
	return "signoz.meter.metric.datapoint.size"
}

// Unit returns the semantic convention unit of the instrument
func (MetricDatapointSize) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (MetricDatapointSize) Description() string {
	return "The size of metric data points observed."
}

// Add adds incr to the existing count for attrs.
//
// The connectorId is the identifier of the meter connector instance that
// produced the metric.
//
// All additional attrs passed are included in the recorded value.
//
// The value is always `0`.
func (m MetricDatapointSize) Add(
	ctx context.Context,
	incr int64,
	connectorId string,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("connector_id", connectorId),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("connector_id", connectorId),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// The value is always `0`.
func (m MetricDatapointSize) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (MetricDatapointSize) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (MetricDatapointSize) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (MetricDatapointSize) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// MetricDatapointSizeObservable is an instrument used to record metric values
// conforming to the "signoz.meter.metric.datapoint.size" semantic conventions.
// It represents the size of metric data points observed.
type MetricDatapointSizeObservable struct {
	metric.Int64ObservableCounter
}

var newMetricDatapointSizeObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The size of metric data points observed."),
	metric.WithUnit("By"),
}

// NewMetricDatapointSizeObservable returns a new MetricDatapointSizeObservable
// instrument.
func NewMetricDatapointSizeObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (MetricDatapointSizeObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return MetricDatapointSizeObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newMetricDatapointSizeObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newMetricDatapointSizeObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz.meter.metric.datapoint.size",
		opt...,
	)
	if err != nil {
		return MetricDatapointSizeObservable{noop.Int64ObservableCounter{}}, err
	}
	return MetricDatapointSizeObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m MetricDatapointSizeObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (MetricDatapointSizeObservable) Name() string {
	return "signoz.meter.metric.datapoint.size"
}

// Unit returns the semantic convention unit of the instrument
func (MetricDatapointSizeObservable) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (MetricDatapointSizeObservable) Description() string {
	return "The size of metric data points observed."
}

// AttrConnectorID returns a required attribute for the "connector_id" semantic
// convention. It represents the identifier of the meter connector instance that
// produced the metric.
func (MetricDatapointSizeObservable) AttrConnectorID(val string) attribute.KeyValue {
	return attribute.String("connector_id", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (MetricDatapointSizeObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (MetricDatapointSizeObservable) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (MetricDatapointSizeObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// SpanCount is an instrument used to record metric values conforming to the
// "signoz.meter.span.count" semantic conventions. It represents the number of
// spans observed.
type SpanCount struct {
	metric.Int64Counter
}

var newSpanCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of spans observed."),
	metric.WithUnit("1"),
}

// NewSpanCount returns a new SpanCount instrument.
func NewSpanCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SpanCount, error) {
	// Check if the meter is nil.
	if m == nil {
		return SpanCount{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newSpanCountOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newSpanCountOpts...)
	}

	i, err := m.Int64Counter(
		"signoz.meter.span.count",
		opt...,
	)
	if err != nil {
		return SpanCount{noop.Int64Counter{}}, err
	}
	return SpanCount{i}, nil
}

// Inst returns the underlying metric instrument.
func (m SpanCount) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (SpanCount) Name() string {
	return "signoz.meter.span.count"
}

// Unit returns the semantic convention unit of the instrument
func (SpanCount) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (SpanCount) Description() string {
	return "The number of spans observed."
}

// Add adds incr to the existing count for attrs.
//
// The connectorId is the identifier of the meter connector instance that
// produced the metric.
//
// All additional attrs passed are included in the recorded value.
//
// Each span counts as one.
func (m SpanCount) Add(
	ctx context.Context,
	incr int64,
	connectorId string,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("connector_id", connectorId),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("connector_id", connectorId),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// Each span counts as one.
func (m SpanCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (SpanCount) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (SpanCount) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (SpanCount) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// SpanCountObservable is an instrument used to record metric values conforming
// to the "signoz.meter.span.count" semantic conventions. It represents the
// number of spans observed.
type SpanCountObservable struct {
	metric.Int64ObservableCounter
}

var newSpanCountObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of spans observed."),
	metric.WithUnit("1"),
}

// NewSpanCountObservable returns a new SpanCountObservable instrument.
func NewSpanCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (SpanCountObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return SpanCountObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newSpanCountObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newSpanCountObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz.meter.span.count",
		opt...,
	)
	if err != nil {
		return SpanCountObservable{noop.Int64ObservableCounter{}}, err
	}
	return SpanCountObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m SpanCountObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (SpanCountObservable) Name() string {
	return "signoz.meter.span.count"
}

// Unit returns the semantic convention unit of the instrument
func (SpanCountObservable) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (SpanCountObservable) Description() string {
	return "The number of spans observed."
}

// AttrConnectorID returns a required attribute for the "connector_id" semantic
// convention. It represents the identifier of the meter connector instance that
// produced the metric.
func (SpanCountObservable) AttrConnectorID(val string) attribute.KeyValue {
	return attribute.String("connector_id", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (SpanCountObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (SpanCountObservable) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (SpanCountObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// SpanSize is an instrument used to record metric values conforming to the
// "signoz.meter.span.size" semantic conventions. It represents the size of spans
// observed.
type SpanSize struct {
	metric.Int64Counter
}

var newSpanSizeOpts = []metric.Int64CounterOption{
	metric.WithDescription("The size of spans observed."),
	metric.WithUnit("By"),
}

// NewSpanSize returns a new SpanSize instrument.
func NewSpanSize(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (SpanSize, error) {
	// Check if the meter is nil.
	if m == nil {
		return SpanSize{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newSpanSizeOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newSpanSizeOpts...)
	}

	i, err := m.Int64Counter(
		"signoz.meter.span.size",
		opt...,
	)
	if err != nil {
		return SpanSize{noop.Int64Counter{}}, err
	}
	return SpanSize{i}, nil
}

// Inst returns the underlying metric instrument.
func (m SpanSize) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (SpanSize) Name() string {
	return "signoz.meter.span.size"
}

// Unit returns the semantic convention unit of the instrument
func (SpanSize) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (SpanSize) Description() string {
	return "The size of spans observed."
}

// Add adds incr to the existing count for attrs.
//
// The connectorId is the identifier of the meter connector instance that
// produced the metric.
//
// All additional attrs passed are included in the recorded value.
//
// The size of a span is the size of its fields, its resource attributes,
// its attributes, its events and its links. `signoz.workspace.*` and
// `signoz.gen_ai.*` attributes are not counted.
func (m SpanSize) Add(
	ctx context.Context,
	incr int64,
	connectorId string,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("connector_id", connectorId),
		))
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(
		*o,
		metric.WithAttributes(
			append(
				attrs[:len(attrs):len(attrs)],
				attribute.String("connector_id", connectorId),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// The size of a span is the size of its fields, its resource attributes,
// its attributes, its events and its links. `signoz.workspace.*` and
// `signoz.gen_ai.*` attributes are not counted.
func (m SpanSize) AddSet(ctx context.Context, incr int64, set attribute.Set) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Int64Counter.Add(ctx, incr)
		return
	}

	o := metricpool.AddOptions()
	defer metricpool.PutAddOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Int64Counter.Add(ctx, incr, *o...)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (SpanSize) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (SpanSize) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (SpanSize) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// SpanSizeObservable is an instrument used to record metric values conforming to
// the "signoz.meter.span.size" semantic conventions. It represents the size of
// spans observed.
type SpanSizeObservable struct {
	metric.Int64ObservableCounter
}

var newSpanSizeObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The size of spans observed."),
	metric.WithUnit("By"),
}

// NewSpanSizeObservable returns a new SpanSizeObservable instrument.
func NewSpanSizeObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (SpanSizeObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return SpanSizeObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newSpanSizeObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newSpanSizeObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz.meter.span.size",
		opt...,
	)
	if err != nil {
		return SpanSizeObservable{noop.Int64ObservableCounter{}}, err
	}
	return SpanSizeObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m SpanSizeObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (SpanSizeObservable) Name() string {
	return "signoz.meter.span.size"
}

// Unit returns the semantic convention unit of the instrument
func (SpanSizeObservable) Unit() string {
	return "By"
}

// Description returns the semantic convention description of the instrument
func (SpanSizeObservable) Description() string {
	return "The size of spans observed."
}

// AttrConnectorID returns a required attribute for the "connector_id" semantic
// convention. It represents the identifier of the meter connector instance that
// produced the metric.
func (SpanSizeObservable) AttrConnectorID(val string) attribute.KeyValue {
	return attribute.String("connector_id", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (SpanSizeObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrHostName returns an optional attribute for the "host.name" semantic
// convention. It represents the name of the host. On UNIX systems, it may
// contain what the hostname command returns, or the fully qualified hostname, or
// another name specified by the user.
func (SpanSizeObservable) AttrHostName(val string) attribute.KeyValue {
	return attribute.String("host.name", val)
}

// AttrServiceName returns an optional attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (SpanSizeObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}
