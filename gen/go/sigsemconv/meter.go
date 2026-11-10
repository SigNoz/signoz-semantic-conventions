// Code generated from semantic convention specification. DO NOT EDIT.

package sigsemconv

import "go.opentelemetry.io/otel/attribute"

// Area: meter
const (
	// ConnectorIDKey is the attribute Key conforming to the "connector_id" semantic
	// conventions. It represents the identifier of the meter connector instance
	// that produced the metric.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "0b8f4a6e-2c1d-4e5f-9a7b-3c6d8e1f2a4b"
	// Note: A UUID assigned when the connector instance is created.
	ConnectorIDKey = attribute.Key("connector_id")

	// SigNozIOWorkloadKey is the attribute Key conforming to the
	// "signoz.io/workload" semantic conventions. It represents the SigNoz workload
	// a Kubernetes node is dedicated to.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "clickhouse"
	// Note: The value of the node's `signoz.io/workload` label.
	//
	// Known values: `clickhouse`, `redpanda`, `consumers`, `zookeeper`,
	// `gateway`, `signoz`.
	SigNozIOWorkloadKey = attribute.Key("signoz.io/workload")
)

// ConnectorID returns an attribute KeyValue conforming to the "connector_id"
// semantic conventions. It represents the identifier of the meter connector
// instance that produced the metric.
func ConnectorID(val string) attribute.KeyValue {
	return ConnectorIDKey.String(val)
}

// SigNozIOWorkload returns an attribute KeyValue conforming to the
// "signoz.io/workload" semantic conventions. It represents the SigNoz workload a
// Kubernetes node is dedicated to.
func SigNozIOWorkload(val string) attribute.KeyValue {
	return SigNozIOWorkloadKey.String(val)
}
