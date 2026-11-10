// Code generated from semantic convention specification. DO NOT EDIT.

// Package spanmetricsconv provides types and functionality for the SigNoz
// semantic conventions in the "spanmetrics" area.
package spanmetricsconv

import (
	"context"

	"github.com/SigNoz/signoz-semantic-conventions/gen/go/sigsemconv/internal/metricpool"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
)

// SpanKindAttr is an attribute conforming to the span.kind semantic conventions.
// It represents the kind of the aggregated spans.
type SpanKindAttr string

var (
	// SpanKindUnspecified is the span kind is unspecified.
	SpanKindUnspecified SpanKindAttr = "SPAN_KIND_UNSPECIFIED"
	// SpanKindInternal is an internal operation.
	SpanKindInternal SpanKindAttr = "SPAN_KIND_INTERNAL"
	// SpanKindServer is the server side of a remote call.
	SpanKindServer SpanKindAttr = "SPAN_KIND_SERVER"
	// SpanKindClient is the client side of a remote call.
	SpanKindClient SpanKindAttr = "SPAN_KIND_CLIENT"
	// SpanKindProducer is the producer of an asynchronous message.
	SpanKindProducer SpanKindAttr = "SPAN_KIND_PRODUCER"
	// SpanKindConsumer is the consumer of an asynchronous message.
	SpanKindConsumer SpanKindAttr = "SPAN_KIND_CONSUMER"
)

// StatusCodeAttr is an attribute conforming to the status.code semantic
// conventions. It represents the status code of the aggregated spans.
type StatusCodeAttr string

var (
	// StatusCodeUnset is the span status is unset.
	StatusCodeUnset StatusCodeAttr = "STATUS_CODE_UNSET"
	// StatusCodeOk is the span completed successfully.
	StatusCodeOk StatusCodeAttr = "STATUS_CODE_OK"
	// StatusCodeError is the span completed with an error.
	StatusCodeError StatusCodeAttr = "STATUS_CODE_ERROR"
)

// CallsTotal is an instrument used to record metric values conforming to the
// "signoz_calls_total" semantic conventions. It represents the number of spans
// observed.
type CallsTotal struct {
	metric.Int64Counter
}

var newCallsTotalOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of spans observed."),
	metric.WithUnit(""),
}

// NewCallsTotal returns a new CallsTotal instrument.
func NewCallsTotal(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (CallsTotal, error) {
	// Check if the meter is nil.
	if m == nil {
		return CallsTotal{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newCallsTotalOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newCallsTotalOpts...)
	}

	i, err := m.Int64Counter(
		"signoz_calls_total",
		opt...,
	)
	if err != nil {
		return CallsTotal{noop.Int64Counter{}}, err
	}
	return CallsTotal{i}, nil
}

// Inst returns the underlying metric instrument.
func (m CallsTotal) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (CallsTotal) Name() string {
	return "signoz_calls_total"
}

// Unit returns the semantic convention unit of the instrument
func (CallsTotal) Unit() string {
	return ""
}

// Description returns the semantic convention description of the instrument
func (CallsTotal) Description() string {
	return "The number of spans observed."
}

// Add adds incr to the existing count for attrs.
//
// The operation is the name of the aggregated spans.
//
// The resourceSigNozCollectorId is a copy of `signoz.collector.id`.
//
// The serviceName is the logical name of the service.
//
// The collectorId is the identifier of the collector instance that produced the
// metric.
//
// The spanKind is the kind of the aggregated spans.
//
// The statusCode is the status code of the aggregated spans.
//
// All additional attrs passed are included in the recorded value.
func (m CallsTotal) Add(
	ctx context.Context,
	incr int64,
	operation string,
	resourceSigNozCollectorId string,
	serviceName string,
	collectorId string,
	spanKind SpanKindAttr,
	statusCode StatusCodeAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("operation", operation),
			attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
			attribute.String("service.name", serviceName),
			attribute.String("signoz.collector.id", collectorId),
			attribute.String("span.kind", string(spanKind)),
			attribute.String("status.code", string(statusCode)),
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
				attribute.String("operation", operation),
				attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
				attribute.String("service.name", serviceName),
				attribute.String("signoz.collector.id", collectorId),
				attribute.String("span.kind", string(spanKind)),
				attribute.String("status.code", string(statusCode)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
func (m CallsTotal) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

// AttrHTTPStatusCode returns an optional attribute for the "http.status_code"
// semantic convention. It represents the deprecated, use
// `http.response.status_code` instead.
func (CallsTotal) AttrHTTPStatusCode(val int) attribute.KeyValue {
	return attribute.Int("http.status_code", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (CallsTotal) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (CallsTotal) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (CallsTotal) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (CallsTotal) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// CallsTotalObservable is an instrument used to record metric values conforming
// to the "signoz_calls_total" semantic conventions. It represents the number of
// spans observed.
type CallsTotalObservable struct {
	metric.Int64ObservableCounter
}

var newCallsTotalObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of spans observed."),
	metric.WithUnit(""),
}

// NewCallsTotalObservable returns a new CallsTotalObservable instrument.
func NewCallsTotalObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (CallsTotalObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return CallsTotalObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newCallsTotalObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newCallsTotalObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz_calls_total",
		opt...,
	)
	if err != nil {
		return CallsTotalObservable{noop.Int64ObservableCounter{}}, err
	}
	return CallsTotalObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m CallsTotalObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (CallsTotalObservable) Name() string {
	return "signoz_calls_total"
}

// Unit returns the semantic convention unit of the instrument
func (CallsTotalObservable) Unit() string {
	return ""
}

// Description returns the semantic convention description of the instrument
func (CallsTotalObservable) Description() string {
	return "The number of spans observed."
}

// AttrOperation returns a required attribute for the "operation" semantic
// convention. It represents the name of the aggregated spans.
func (CallsTotalObservable) AttrOperation(val string) attribute.KeyValue {
	return attribute.String("operation", val)
}

// AttrResourceSigNozCollectorID returns a required attribute for the
// "resource_signoz.collector.id" semantic convention. It represents a copy of
// `signoz.collector.id`.
func (CallsTotalObservable) AttrResourceSigNozCollectorID(val string) attribute.KeyValue {
	return attribute.String("resource_signoz.collector.id", val)
}

// AttrServiceName returns a required attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (CallsTotalObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// AttrCollectorID returns a required attribute for the "signoz.collector.id"
// semantic convention. It represents the identifier of the collector instance
// that produced the metric.
func (CallsTotalObservable) AttrCollectorID(val string) attribute.KeyValue {
	return attribute.String("signoz.collector.id", val)
}

// AttrSpanKind returns a required attribute for the "span.kind" semantic
// convention. It represents the kind of the aggregated spans.
func (CallsTotalObservable) AttrSpanKind(val SpanKindAttr) attribute.KeyValue {
	return attribute.String("span.kind", string(val))
}

// AttrStatusCode returns a required attribute for the "status.code" semantic
// convention. It represents the status code of the aggregated spans.
func (CallsTotalObservable) AttrStatusCode(val StatusCodeAttr) attribute.KeyValue {
	return attribute.String("status.code", string(val))
}

// AttrHTTPStatusCode returns an optional attribute for the "http.status_code"
// semantic convention. It represents the deprecated, use
// `http.response.status_code` instead.
func (CallsTotalObservable) AttrHTTPStatusCode(val int) attribute.KeyValue {
	return attribute.Int("http.status_code", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (CallsTotalObservable) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (CallsTotalObservable) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (CallsTotalObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (CallsTotalObservable) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// DBLatencyCount is an instrument used to record metric values conforming to the
// "signoz_db_latency_count" semantic conventions. It represents the number of
// database calls observed.
type DBLatencyCount struct {
	metric.Int64Counter
}

var newDBLatencyCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of database calls observed."),
	metric.WithUnit("1"),
}

// NewDBLatencyCount returns a new DBLatencyCount instrument.
func NewDBLatencyCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (DBLatencyCount, error) {
	// Check if the meter is nil.
	if m == nil {
		return DBLatencyCount{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newDBLatencyCountOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newDBLatencyCountOpts...)
	}

	i, err := m.Int64Counter(
		"signoz_db_latency_count",
		opt...,
	)
	if err != nil {
		return DBLatencyCount{noop.Int64Counter{}}, err
	}
	return DBLatencyCount{i}, nil
}

// Inst returns the underlying metric instrument.
func (m DBLatencyCount) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (DBLatencyCount) Name() string {
	return "signoz_db_latency_count"
}

// Unit returns the semantic convention unit of the instrument
func (DBLatencyCount) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (DBLatencyCount) Description() string {
	return "The number of database calls observed."
}

// Add adds incr to the existing count for attrs.
//
// The dbSystem is the deprecated, use `db.system.name` instead.
//
// The resourceSigNozCollectorId is a copy of `signoz.collector.id`.
//
// The serviceName is the logical name of the service.
//
// The collectorId is the identifier of the collector instance that produced the
// metric.
//
// The statusCode is the status code of the aggregated spans.
//
// All additional attrs passed are included in the recorded value.
//
// A database call is a span that carries `db.system` and is not of kind
// server.
func (m DBLatencyCount) Add(
	ctx context.Context,
	incr int64,
	dbSystem string,
	resourceSigNozCollectorId string,
	serviceName string,
	collectorId string,
	statusCode StatusCodeAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("db.system", dbSystem),
			attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
			attribute.String("service.name", serviceName),
			attribute.String("signoz.collector.id", collectorId),
			attribute.String("status.code", string(statusCode)),
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
				attribute.String("db.system", dbSystem),
				attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
				attribute.String("service.name", serviceName),
				attribute.String("signoz.collector.id", collectorId),
				attribute.String("status.code", string(statusCode)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// A database call is a span that carries `db.system` and is not of kind
// server.
func (m DBLatencyCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

// AttrDBName returns an optional attribute for the "db.name" semantic
// convention. It represents the deprecated, use `db.namespace` instead.
func (DBLatencyCount) AttrDBName(val string) attribute.KeyValue {
	return attribute.String("db.name", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (DBLatencyCount) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (DBLatencyCount) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (DBLatencyCount) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (DBLatencyCount) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// DBLatencyCountObservable is an instrument used to record metric values
// conforming to the "signoz_db_latency_count" semantic conventions. It
// represents the number of database calls observed.
type DBLatencyCountObservable struct {
	metric.Int64ObservableCounter
}

var newDBLatencyCountObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of database calls observed."),
	metric.WithUnit("1"),
}

// NewDBLatencyCountObservable returns a new DBLatencyCountObservable instrument.
func NewDBLatencyCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (DBLatencyCountObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return DBLatencyCountObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newDBLatencyCountObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newDBLatencyCountObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz_db_latency_count",
		opt...,
	)
	if err != nil {
		return DBLatencyCountObservable{noop.Int64ObservableCounter{}}, err
	}
	return DBLatencyCountObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m DBLatencyCountObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (DBLatencyCountObservable) Name() string {
	return "signoz_db_latency_count"
}

// Unit returns the semantic convention unit of the instrument
func (DBLatencyCountObservable) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (DBLatencyCountObservable) Description() string {
	return "The number of database calls observed."
}

// AttrDBSystem returns a required attribute for the "db.system" semantic
// convention. It represents the deprecated, use `db.system.name` instead.
func (DBLatencyCountObservable) AttrDBSystem(val string) attribute.KeyValue {
	return attribute.String("db.system", val)
}

// AttrResourceSigNozCollectorID returns a required attribute for the
// "resource_signoz.collector.id" semantic convention. It represents a copy of
// `signoz.collector.id`.
func (DBLatencyCountObservable) AttrResourceSigNozCollectorID(val string) attribute.KeyValue {
	return attribute.String("resource_signoz.collector.id", val)
}

// AttrServiceName returns a required attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (DBLatencyCountObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// AttrCollectorID returns a required attribute for the "signoz.collector.id"
// semantic convention. It represents the identifier of the collector instance
// that produced the metric.
func (DBLatencyCountObservable) AttrCollectorID(val string) attribute.KeyValue {
	return attribute.String("signoz.collector.id", val)
}

// AttrStatusCode returns a required attribute for the "status.code" semantic
// convention. It represents the status code of the aggregated spans.
func (DBLatencyCountObservable) AttrStatusCode(val StatusCodeAttr) attribute.KeyValue {
	return attribute.String("status.code", string(val))
}

// AttrDBName returns an optional attribute for the "db.name" semantic
// convention. It represents the deprecated, use `db.namespace` instead.
func (DBLatencyCountObservable) AttrDBName(val string) attribute.KeyValue {
	return attribute.String("db.name", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (DBLatencyCountObservable) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (DBLatencyCountObservable) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (DBLatencyCountObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (DBLatencyCountObservable) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// DBLatencySum is an instrument used to record metric values conforming to the
// "signoz_db_latency_sum" semantic conventions. It represents the total duration
// of database calls observed.
type DBLatencySum struct {
	metric.Int64Counter
}

var newDBLatencySumOpts = []metric.Int64CounterOption{
	metric.WithDescription("The total duration of database calls observed."),
	metric.WithUnit("1"),
}

// NewDBLatencySum returns a new DBLatencySum instrument.
func NewDBLatencySum(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (DBLatencySum, error) {
	// Check if the meter is nil.
	if m == nil {
		return DBLatencySum{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newDBLatencySumOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newDBLatencySumOpts...)
	}

	i, err := m.Int64Counter(
		"signoz_db_latency_sum",
		opt...,
	)
	if err != nil {
		return DBLatencySum{noop.Int64Counter{}}, err
	}
	return DBLatencySum{i}, nil
}

// Inst returns the underlying metric instrument.
func (m DBLatencySum) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (DBLatencySum) Name() string {
	return "signoz_db_latency_sum"
}

// Unit returns the semantic convention unit of the instrument
func (DBLatencySum) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (DBLatencySum) Description() string {
	return "The total duration of database calls observed."
}

// Add adds incr to the existing count for attrs.
//
// The dbSystem is the deprecated, use `db.system.name` instead.
//
// The resourceSigNozCollectorId is a copy of `signoz.collector.id`.
//
// The serviceName is the logical name of the service.
//
// The collectorId is the identifier of the collector instance that produced the
// metric.
//
// The statusCode is the status code of the aggregated spans.
//
// All additional attrs passed are included in the recorded value.
//
// Values are in milliseconds. A database call is a span that carries
// `db.system` and is not of kind server.
func (m DBLatencySum) Add(
	ctx context.Context,
	incr int64,
	dbSystem string,
	resourceSigNozCollectorId string,
	serviceName string,
	collectorId string,
	statusCode StatusCodeAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("db.system", dbSystem),
			attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
			attribute.String("service.name", serviceName),
			attribute.String("signoz.collector.id", collectorId),
			attribute.String("status.code", string(statusCode)),
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
				attribute.String("db.system", dbSystem),
				attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
				attribute.String("service.name", serviceName),
				attribute.String("signoz.collector.id", collectorId),
				attribute.String("status.code", string(statusCode)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// Values are in milliseconds. A database call is a span that carries
// `db.system` and is not of kind server.
func (m DBLatencySum) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

// AttrDBName returns an optional attribute for the "db.name" semantic
// convention. It represents the deprecated, use `db.namespace` instead.
func (DBLatencySum) AttrDBName(val string) attribute.KeyValue {
	return attribute.String("db.name", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (DBLatencySum) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (DBLatencySum) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (DBLatencySum) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (DBLatencySum) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// DBLatencySumObservable is an instrument used to record metric values
// conforming to the "signoz_db_latency_sum" semantic conventions. It represents
// the total duration of database calls observed.
type DBLatencySumObservable struct {
	metric.Int64ObservableCounter
}

var newDBLatencySumObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The total duration of database calls observed."),
	metric.WithUnit("1"),
}

// NewDBLatencySumObservable returns a new DBLatencySumObservable instrument.
func NewDBLatencySumObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (DBLatencySumObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return DBLatencySumObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newDBLatencySumObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newDBLatencySumObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz_db_latency_sum",
		opt...,
	)
	if err != nil {
		return DBLatencySumObservable{noop.Int64ObservableCounter{}}, err
	}
	return DBLatencySumObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m DBLatencySumObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (DBLatencySumObservable) Name() string {
	return "signoz_db_latency_sum"
}

// Unit returns the semantic convention unit of the instrument
func (DBLatencySumObservable) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (DBLatencySumObservable) Description() string {
	return "The total duration of database calls observed."
}

// AttrDBSystem returns a required attribute for the "db.system" semantic
// convention. It represents the deprecated, use `db.system.name` instead.
func (DBLatencySumObservable) AttrDBSystem(val string) attribute.KeyValue {
	return attribute.String("db.system", val)
}

// AttrResourceSigNozCollectorID returns a required attribute for the
// "resource_signoz.collector.id" semantic convention. It represents a copy of
// `signoz.collector.id`.
func (DBLatencySumObservable) AttrResourceSigNozCollectorID(val string) attribute.KeyValue {
	return attribute.String("resource_signoz.collector.id", val)
}

// AttrServiceName returns a required attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (DBLatencySumObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// AttrCollectorID returns a required attribute for the "signoz.collector.id"
// semantic convention. It represents the identifier of the collector instance
// that produced the metric.
func (DBLatencySumObservable) AttrCollectorID(val string) attribute.KeyValue {
	return attribute.String("signoz.collector.id", val)
}

// AttrStatusCode returns a required attribute for the "status.code" semantic
// convention. It represents the status code of the aggregated spans.
func (DBLatencySumObservable) AttrStatusCode(val StatusCodeAttr) attribute.KeyValue {
	return attribute.String("status.code", string(val))
}

// AttrDBName returns an optional attribute for the "db.name" semantic
// convention. It represents the deprecated, use `db.namespace` instead.
func (DBLatencySumObservable) AttrDBName(val string) attribute.KeyValue {
	return attribute.String("db.name", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (DBLatencySumObservable) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (DBLatencySumObservable) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (DBLatencySumObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (DBLatencySumObservable) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// ExternalCallLatencyCount is an instrument used to record metric values
// conforming to the "signoz_external_call_latency_count" semantic conventions.
// It represents the number of external calls observed.
type ExternalCallLatencyCount struct {
	metric.Int64Counter
}

var newExternalCallLatencyCountOpts = []metric.Int64CounterOption{
	metric.WithDescription("The number of external calls observed."),
	metric.WithUnit("1"),
}

// NewExternalCallLatencyCount returns a new ExternalCallLatencyCount instrument.
func NewExternalCallLatencyCount(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ExternalCallLatencyCount, error) {
	// Check if the meter is nil.
	if m == nil {
		return ExternalCallLatencyCount{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newExternalCallLatencyCountOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newExternalCallLatencyCountOpts...)
	}

	i, err := m.Int64Counter(
		"signoz_external_call_latency_count",
		opt...,
	)
	if err != nil {
		return ExternalCallLatencyCount{noop.Int64Counter{}}, err
	}
	return ExternalCallLatencyCount{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ExternalCallLatencyCount) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (ExternalCallLatencyCount) Name() string {
	return "signoz_external_call_latency_count"
}

// Unit returns the semantic convention unit of the instrument
func (ExternalCallLatencyCount) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (ExternalCallLatencyCount) Description() string {
	return "The number of external calls observed."
}

// Add adds incr to the existing count for attrs.
//
// The address is the remote endpoint of an external call.
//
// The resourceSigNozCollectorId is a copy of `signoz.collector.id`.
//
// The serviceName is the logical name of the service.
//
// The collectorId is the identifier of the collector instance that produced the
// metric.
//
// The statusCode is the status code of the aggregated spans.
//
// All additional attrs passed are included in the recorded value.
//
// An external call is a client span with a known remote endpoint.
func (m ExternalCallLatencyCount) Add(
	ctx context.Context,
	incr int64,
	address string,
	resourceSigNozCollectorId string,
	serviceName string,
	collectorId string,
	statusCode StatusCodeAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("address", address),
			attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
			attribute.String("service.name", serviceName),
			attribute.String("signoz.collector.id", collectorId),
			attribute.String("status.code", string(statusCode)),
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
				attribute.String("address", address),
				attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
				attribute.String("service.name", serviceName),
				attribute.String("signoz.collector.id", collectorId),
				attribute.String("status.code", string(statusCode)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// An external call is a client span with a known remote endpoint.
func (m ExternalCallLatencyCount) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

// AttrHTTPStatusCode returns an optional attribute for the "http.status_code"
// semantic convention. It represents the deprecated, use
// `http.response.status_code` instead.
func (ExternalCallLatencyCount) AttrHTTPStatusCode(val int) attribute.KeyValue {
	return attribute.Int("http.status_code", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (ExternalCallLatencyCount) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (ExternalCallLatencyCount) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (ExternalCallLatencyCount) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (ExternalCallLatencyCount) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// ExternalCallLatencyCountObservable is an instrument used to record metric
// values conforming to the "signoz_external_call_latency_count" semantic
// conventions. It represents the number of external calls observed.
type ExternalCallLatencyCountObservable struct {
	metric.Int64ObservableCounter
}

var newExternalCallLatencyCountObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The number of external calls observed."),
	metric.WithUnit("1"),
}

// NewExternalCallLatencyCountObservable returns a new
// ExternalCallLatencyCountObservable instrument.
func NewExternalCallLatencyCountObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ExternalCallLatencyCountObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return ExternalCallLatencyCountObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newExternalCallLatencyCountObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newExternalCallLatencyCountObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz_external_call_latency_count",
		opt...,
	)
	if err != nil {
		return ExternalCallLatencyCountObservable{noop.Int64ObservableCounter{}}, err
	}
	return ExternalCallLatencyCountObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ExternalCallLatencyCountObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (ExternalCallLatencyCountObservable) Name() string {
	return "signoz_external_call_latency_count"
}

// Unit returns the semantic convention unit of the instrument
func (ExternalCallLatencyCountObservable) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (ExternalCallLatencyCountObservable) Description() string {
	return "The number of external calls observed."
}

// AttrAddress returns a required attribute for the "address" semantic
// convention. It represents the remote endpoint of an external call.
func (ExternalCallLatencyCountObservable) AttrAddress(val string) attribute.KeyValue {
	return attribute.String("address", val)
}

// AttrResourceSigNozCollectorID returns a required attribute for the
// "resource_signoz.collector.id" semantic convention. It represents a copy of
// `signoz.collector.id`.
func (ExternalCallLatencyCountObservable) AttrResourceSigNozCollectorID(val string) attribute.KeyValue {
	return attribute.String("resource_signoz.collector.id", val)
}

// AttrServiceName returns a required attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (ExternalCallLatencyCountObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// AttrCollectorID returns a required attribute for the "signoz.collector.id"
// semantic convention. It represents the identifier of the collector instance
// that produced the metric.
func (ExternalCallLatencyCountObservable) AttrCollectorID(val string) attribute.KeyValue {
	return attribute.String("signoz.collector.id", val)
}

// AttrStatusCode returns a required attribute for the "status.code" semantic
// convention. It represents the status code of the aggregated spans.
func (ExternalCallLatencyCountObservable) AttrStatusCode(val StatusCodeAttr) attribute.KeyValue {
	return attribute.String("status.code", string(val))
}

// AttrHTTPStatusCode returns an optional attribute for the "http.status_code"
// semantic convention. It represents the deprecated, use
// `http.response.status_code` instead.
func (ExternalCallLatencyCountObservable) AttrHTTPStatusCode(val int) attribute.KeyValue {
	return attribute.Int("http.status_code", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (ExternalCallLatencyCountObservable) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (ExternalCallLatencyCountObservable) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (ExternalCallLatencyCountObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (ExternalCallLatencyCountObservable) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// ExternalCallLatencySum is an instrument used to record metric values
// conforming to the "signoz_external_call_latency_sum" semantic conventions. It
// represents the total duration of external calls observed.
type ExternalCallLatencySum struct {
	metric.Int64Counter
}

var newExternalCallLatencySumOpts = []metric.Int64CounterOption{
	metric.WithDescription("The total duration of external calls observed."),
	metric.WithUnit("1"),
}

// NewExternalCallLatencySum returns a new ExternalCallLatencySum instrument.
func NewExternalCallLatencySum(
	m metric.Meter,
	opt ...metric.Int64CounterOption,
) (ExternalCallLatencySum, error) {
	// Check if the meter is nil.
	if m == nil {
		return ExternalCallLatencySum{noop.Int64Counter{}}, nil
	}

	if len(opt) == 0 {
		opt = newExternalCallLatencySumOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newExternalCallLatencySumOpts...)
	}

	i, err := m.Int64Counter(
		"signoz_external_call_latency_sum",
		opt...,
	)
	if err != nil {
		return ExternalCallLatencySum{noop.Int64Counter{}}, err
	}
	return ExternalCallLatencySum{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ExternalCallLatencySum) Inst() metric.Int64Counter {
	return m.Int64Counter
}

// Name returns the semantic convention name of the instrument.
func (ExternalCallLatencySum) Name() string {
	return "signoz_external_call_latency_sum"
}

// Unit returns the semantic convention unit of the instrument
func (ExternalCallLatencySum) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (ExternalCallLatencySum) Description() string {
	return "The total duration of external calls observed."
}

// Add adds incr to the existing count for attrs.
//
// The address is the remote endpoint of an external call.
//
// The resourceSigNozCollectorId is a copy of `signoz.collector.id`.
//
// The serviceName is the logical name of the service.
//
// The collectorId is the identifier of the collector instance that produced the
// metric.
//
// The statusCode is the status code of the aggregated spans.
//
// All additional attrs passed are included in the recorded value.
//
// Values are in milliseconds. An external call is a client span with a
// known remote endpoint.
func (m ExternalCallLatencySum) Add(
	ctx context.Context,
	incr int64,
	address string,
	resourceSigNozCollectorId string,
	serviceName string,
	collectorId string,
	statusCode StatusCodeAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Int64Counter.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Int64Counter.Add(ctx, incr, metric.WithAttributes(
			attribute.String("address", address),
			attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
			attribute.String("service.name", serviceName),
			attribute.String("signoz.collector.id", collectorId),
			attribute.String("status.code", string(statusCode)),
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
				attribute.String("address", address),
				attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
				attribute.String("service.name", serviceName),
				attribute.String("signoz.collector.id", collectorId),
				attribute.String("status.code", string(statusCode)),
			)...,
		),
	)

	m.Int64Counter.Add(ctx, incr, *o...)
}

// AddSet adds incr to the existing count for set.
//
// Values are in milliseconds. An external call is a client span with a
// known remote endpoint.
func (m ExternalCallLatencySum) AddSet(ctx context.Context, incr int64, set attribute.Set) {
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

// AttrHTTPStatusCode returns an optional attribute for the "http.status_code"
// semantic convention. It represents the deprecated, use
// `http.response.status_code` instead.
func (ExternalCallLatencySum) AttrHTTPStatusCode(val int) attribute.KeyValue {
	return attribute.Int("http.status_code", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (ExternalCallLatencySum) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (ExternalCallLatencySum) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (ExternalCallLatencySum) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (ExternalCallLatencySum) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// ExternalCallLatencySumObservable is an instrument used to record metric values
// conforming to the "signoz_external_call_latency_sum" semantic conventions. It
// represents the total duration of external calls observed.
type ExternalCallLatencySumObservable struct {
	metric.Int64ObservableCounter
}

var newExternalCallLatencySumObservableOpts = []metric.Int64ObservableCounterOption{
	metric.WithDescription("The total duration of external calls observed."),
	metric.WithUnit("1"),
}

// NewExternalCallLatencySumObservable returns a new
// ExternalCallLatencySumObservable instrument.
func NewExternalCallLatencySumObservable(
	m metric.Meter,
	opt ...metric.Int64ObservableCounterOption,
) (ExternalCallLatencySumObservable, error) {
	// Check if the meter is nil.
	if m == nil {
		return ExternalCallLatencySumObservable{noop.Int64ObservableCounter{}}, nil
	}

	if len(opt) == 0 {
		opt = newExternalCallLatencySumObservableOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newExternalCallLatencySumObservableOpts...)
	}

	i, err := m.Int64ObservableCounter(
		"signoz_external_call_latency_sum",
		opt...,
	)
	if err != nil {
		return ExternalCallLatencySumObservable{noop.Int64ObservableCounter{}}, err
	}
	return ExternalCallLatencySumObservable{i}, nil
}

// Inst returns the underlying metric instrument.
func (m ExternalCallLatencySumObservable) Inst() metric.Int64ObservableCounter {
	return m.Int64ObservableCounter
}

// Name returns the semantic convention name of the instrument.
func (ExternalCallLatencySumObservable) Name() string {
	return "signoz_external_call_latency_sum"
}

// Unit returns the semantic convention unit of the instrument
func (ExternalCallLatencySumObservable) Unit() string {
	return "1"
}

// Description returns the semantic convention description of the instrument
func (ExternalCallLatencySumObservable) Description() string {
	return "The total duration of external calls observed."
}

// AttrAddress returns a required attribute for the "address" semantic
// convention. It represents the remote endpoint of an external call.
func (ExternalCallLatencySumObservable) AttrAddress(val string) attribute.KeyValue {
	return attribute.String("address", val)
}

// AttrResourceSigNozCollectorID returns a required attribute for the
// "resource_signoz.collector.id" semantic convention. It represents a copy of
// `signoz.collector.id`.
func (ExternalCallLatencySumObservable) AttrResourceSigNozCollectorID(val string) attribute.KeyValue {
	return attribute.String("resource_signoz.collector.id", val)
}

// AttrServiceName returns a required attribute for the "service.name" semantic
// convention. It represents the logical name of the service.
func (ExternalCallLatencySumObservable) AttrServiceName(val string) attribute.KeyValue {
	return attribute.String("service.name", val)
}

// AttrCollectorID returns a required attribute for the "signoz.collector.id"
// semantic convention. It represents the identifier of the collector instance
// that produced the metric.
func (ExternalCallLatencySumObservable) AttrCollectorID(val string) attribute.KeyValue {
	return attribute.String("signoz.collector.id", val)
}

// AttrStatusCode returns a required attribute for the "status.code" semantic
// convention. It represents the status code of the aggregated spans.
func (ExternalCallLatencySumObservable) AttrStatusCode(val StatusCodeAttr) attribute.KeyValue {
	return attribute.String("status.code", string(val))
}

// AttrHTTPStatusCode returns an optional attribute for the "http.status_code"
// semantic convention. It represents the deprecated, use
// `http.response.status_code` instead.
func (ExternalCallLatencySumObservable) AttrHTTPStatusCode(val int) attribute.KeyValue {
	return attribute.Int("http.status_code", val)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (ExternalCallLatencySumObservable) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (ExternalCallLatencySumObservable) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (ExternalCallLatencySumObservable) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (ExternalCallLatencySumObservable) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}

// Latency is an instrument used to record metric values conforming to the
// "signoz_latency" semantic conventions. It represents the duration of spans
// observed.
type Latency struct {
	metric.Float64Histogram
}

var newLatencyOpts = []metric.Float64HistogramOption{
	metric.WithDescription("The duration of spans observed."),
	metric.WithUnit("ms"),
}

// NewLatency returns a new Latency instrument.
func NewLatency(
	m metric.Meter,
	opt ...metric.Float64HistogramOption,
) (Latency, error) {
	// Check if the meter is nil.
	if m == nil {
		return Latency{noop.Float64Histogram{}}, nil
	}

	if len(opt) == 0 {
		opt = newLatencyOpts
	} else {
		opt = append(opt[:len(opt):len(opt)], newLatencyOpts...)
	}

	i, err := m.Float64Histogram(
		"signoz_latency",
		opt...,
	)
	if err != nil {
		return Latency{noop.Float64Histogram{}}, err
	}
	return Latency{i}, nil
}

// Inst returns the underlying metric instrument.
func (m Latency) Inst() metric.Float64Histogram {
	return m.Float64Histogram
}

// Name returns the semantic convention name of the instrument.
func (Latency) Name() string {
	return "signoz_latency"
}

// Unit returns the semantic convention unit of the instrument
func (Latency) Unit() string {
	return "ms"
}

// Description returns the semantic convention description of the instrument
func (Latency) Description() string {
	return "The duration of spans observed."
}

// Record records val to the current distribution for attrs.
//
// The operation is the name of the aggregated spans.
//
// The resourceSigNozCollectorId is a copy of `signoz.collector.id`.
//
// The serviceName is the logical name of the service.
//
// The collectorId is the identifier of the collector instance that produced the
// metric.
//
// The spanKind is the kind of the aggregated spans.
//
// The statusCode is the status code of the aggregated spans.
//
// All additional attrs passed are included in the recorded value.
//
// Recorded as an explicit bucket histogram and, when enabled, also as an
// exponential histogram under the same name. Values are in milliseconds.
func (m Latency) Record(
	ctx context.Context,
	val float64,
	operation string,
	resourceSigNozCollectorId string,
	serviceName string,
	collectorId string,
	spanKind SpanKindAttr,
	statusCode StatusCodeAttr,
	attrs ...attribute.KeyValue,
) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if len(attrs) == 0 {
		m.Float64Histogram.Record(ctx, val, metric.WithAttributes(
			attribute.String("operation", operation),
			attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
			attribute.String("service.name", serviceName),
			attribute.String("signoz.collector.id", collectorId),
			attribute.String("span.kind", string(spanKind)),
			attribute.String("status.code", string(statusCode)),
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
				attribute.String("operation", operation),
				attribute.String("resource_signoz.collector.id", resourceSigNozCollectorId),
				attribute.String("service.name", serviceName),
				attribute.String("signoz.collector.id", collectorId),
				attribute.String("span.kind", string(spanKind)),
				attribute.String("status.code", string(statusCode)),
			)...,
		),
	)

	m.Float64Histogram.Record(ctx, val, *o...)
}

// RecordSet records val to the current distribution for set.
//
// Recorded as an explicit bucket histogram and, when enabled, also as an
// exponential histogram under the same name. Values are in milliseconds.
func (m Latency) RecordSet(ctx context.Context, val float64, set attribute.Set) {
	if !m.Float64Histogram.Enabled(ctx) {
		return
	}
	if set.Len() == 0 {
		m.Float64Histogram.Record(ctx, val)
		return
	}

	o := metricpool.RecordOptions()
	defer metricpool.PutRecordOptions(o)

	*o = append(*o, metric.WithAttributeSet(set))
	m.Float64Histogram.Record(ctx, val, *o...)
}

// AttrResourceDeploymentEnvironment returns an optional attribute for the
// "resource_deployment.environment" semantic convention. It represents the value
// of `deployment.environment` on the span's Resource.
func (Latency) AttrResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("resource_deployment.environment", val)
}

// AttrResourceServiceNamespace returns an optional attribute for the
// "resource_service.namespace" semantic convention. It represents the value of
// `service.namespace` on the span's Resource.
func (Latency) AttrResourceServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("resource_service.namespace", val)
}

// AttrDeploymentEnvironment returns an optional attribute for the
// "deployment.environment" semantic convention. It represents the deprecated,
// use `deployment.environment.name` instead.
func (Latency) AttrDeploymentEnvironment(val string) attribute.KeyValue {
	return attribute.String("deployment.environment", val)
}

// AttrServiceNamespace returns an optional attribute for the "service.namespace"
// semantic convention. It represents a namespace for `service.name`.
func (Latency) AttrServiceNamespace(val string) attribute.KeyValue {
	return attribute.String("service.namespace", val)
}
