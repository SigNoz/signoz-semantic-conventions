// Code generated from semantic convention specification. DO NOT EDIT.

package sigsemconv

import "go.opentelemetry.io/otel/attribute"

// Area: spanmetrics
const (
	// AddressKey is the attribute Key conforming to the "address" semantic
	// conventions. It represents the remote endpoint of an external call.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "payments.internal:8443", "oteldemo.CartService/GetCart"
	// Note: `<rpc.service>/<rpc.method>` for RPC calls, otherwise the remote host
	// as `host` or `host:port`, or the span's `peer.service`.
	AddressKey = attribute.Key("address")

	// OperationKey is the attribute Key conforming to the "operation" semantic
	// conventions. It represents the name of the aggregated spans.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "GET /api/v1/orders/{id}", "overflow_operation"
	// Note: Set to `overflow_operation` when the service has more distinct span
	// names than are tracked.
	OperationKey = attribute.Key("operation")

	// ResourceDeploymentEnvironmentKey is the attribute Key conforming to the
	// "resource_deployment.environment" semantic conventions. It represents the
	// value of `deployment.environment` on the span's Resource.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "production"
	ResourceDeploymentEnvironmentKey = attribute.Key("resource_deployment.environment")

	// ResourceServiceNamespaceKey is the attribute Key conforming to the
	// "resource_service.namespace" semantic conventions. It represents the value of
	// `service.namespace` on the span's Resource.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "shop"
	ResourceServiceNamespaceKey = attribute.Key("resource_service.namespace")

	// ResourceSigNozCollectorIDKey is the attribute Key conforming to the
	// "resource_signoz.collector.id" semantic conventions. It represents a copy of
	// `signoz.collector.id`.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "6f1c2b7e-4d3a-4b8f-9e21-5a0c7d9b3e14"
	ResourceSigNozCollectorIDKey = attribute.Key("resource_signoz.collector.id")

	// SigNozCollectorIDKey is the attribute Key conforming to the
	// "signoz.collector.id" semantic conventions. It represents the identifier of
	// the collector instance that produced the metric.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "6f1c2b7e-4d3a-4b8f-9e21-5a0c7d9b3e14"
	// Note: The collector's `service.instance.id`, or a UUID assigned when the
	// collector starts if it has none.
	SigNozCollectorIDKey = attribute.Key("signoz.collector.id")

	// SpanKindKey is the attribute Key conforming to the "span.kind" semantic
	// conventions. It represents the kind of the aggregated spans.
	//
	// Type: Enum
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "SPAN_KIND_SERVER"
	SpanKindKey = attribute.Key("span.kind")

	// StatusCodeKey is the attribute Key conforming to the "status.code" semantic
	// conventions. It represents the status code of the aggregated spans.
	//
	// Type: Enum
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "STATUS_CODE_ERROR"
	StatusCodeKey = attribute.Key("status.code")
)

// Address returns an attribute KeyValue conforming to the "address" semantic
// conventions. It represents the remote endpoint of an external call.
func Address(val string) attribute.KeyValue {
	return AddressKey.String(val)
}

// Operation returns an attribute KeyValue conforming to the "operation" semantic
// conventions. It represents the name of the aggregated spans.
func Operation(val string) attribute.KeyValue {
	return OperationKey.String(val)
}

// ResourceDeploymentEnvironment returns an attribute KeyValue conforming to the
// "resource_deployment.environment" semantic conventions. It represents the
// value of `deployment.environment` on the span's Resource.
func ResourceDeploymentEnvironment(val string) attribute.KeyValue {
	return ResourceDeploymentEnvironmentKey.String(val)
}

// ResourceServiceNamespace returns an attribute KeyValue conforming to the
// "resource_service.namespace" semantic conventions. It represents the value of
// `service.namespace` on the span's Resource.
func ResourceServiceNamespace(val string) attribute.KeyValue {
	return ResourceServiceNamespaceKey.String(val)
}

// ResourceSigNozCollectorID returns an attribute KeyValue conforming to the
// "resource_signoz.collector.id" semantic conventions. It represents a copy of
// `signoz.collector.id`.
func ResourceSigNozCollectorID(val string) attribute.KeyValue {
	return ResourceSigNozCollectorIDKey.String(val)
}

// SigNozCollectorID returns an attribute KeyValue conforming to the
// "signoz.collector.id" semantic conventions. It represents the identifier of
// the collector instance that produced the metric.
func SigNozCollectorID(val string) attribute.KeyValue {
	return SigNozCollectorIDKey.String(val)
}

// Enum values for span.kind
var (
	// The span kind is unspecified.
	//
	// Stability: development
	SpanKindUnspecified = SpanKindKey.String("SPAN_KIND_UNSPECIFIED")
	// An internal operation.
	//
	// Stability: development
	SpanKindInternal = SpanKindKey.String("SPAN_KIND_INTERNAL")
	// The server side of a remote call.
	//
	// Stability: development
	SpanKindServer = SpanKindKey.String("SPAN_KIND_SERVER")
	// The client side of a remote call.
	//
	// Stability: development
	SpanKindClient = SpanKindKey.String("SPAN_KIND_CLIENT")
	// The producer of an asynchronous message.
	//
	// Stability: development
	SpanKindProducer = SpanKindKey.String("SPAN_KIND_PRODUCER")
	// The consumer of an asynchronous message.
	//
	// Stability: development
	SpanKindConsumer = SpanKindKey.String("SPAN_KIND_CONSUMER")
)

// Enum values for status.code
var (
	// The span status is unset.
	//
	// Stability: development
	StatusCodeUnset = StatusCodeKey.String("STATUS_CODE_UNSET")
	// The span completed successfully.
	//
	// Stability: development
	StatusCodeOk = StatusCodeKey.String("STATUS_CODE_OK")
	// The span completed with an error.
	//
	// Stability: development
	StatusCodeError = StatusCodeKey.String("STATUS_CODE_ERROR")
)
