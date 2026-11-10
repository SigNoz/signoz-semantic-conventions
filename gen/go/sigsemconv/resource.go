// Code generated from semantic convention specification. DO NOT EDIT.

package sigsemconv

import "go.opentelemetry.io/otel/attribute"

// Area: resource
const (
	// SigNozResourceIDKey is the attribute Key conforming to the
	// "signoz.resource.id" semantic conventions. It represents the unique
	// identifier of the resource.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "3f8a2c14-9b7d-4e1f-a6c3-5d8e2f0b1a4c", "*"
	// Note: The id is unique within its kind and organization, not across kinds.
	// Its format depends on the kind; most platform objects use a UUID.
	//
	// The value `*` denotes the collection of all resources of the kind and
	// is used when an operation applies to the collection rather than to a
	// single resource, for example a list operation.
	SigNozResourceIDKey = attribute.Key("signoz.resource.id")

	// SigNozResourceKindKey is the attribute Key conforming to the
	// "signoz.resource.kind" semantic conventions. It represents the kind of the
	// resource.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "ingestion-key"
	// Note: Kind values are lowercase, with words joined by hyphens.
	//
	// Known values include metaresource kinds such as `dashboard` and
	// `alert-rule`; telemetry resource kinds `logs`, `traces`, `metrics`,
	// `meter-metrics`, `logs-field` and `traces-field`; and platform object
	// kinds such as `ingestion-key`, `ingestion-limit`, `license`,
	// `subscription` and `deployment-host`. The set is not final.
	SigNozResourceKindKey = attribute.Key("signoz.resource.kind")

	// SigNozResourceObjectKey is the attribute Key conforming to the
	// "signoz.resource.object" semantic conventions. It represents the object
	// reference of the resource in the authorization model.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples:
	// "telemetryresource:organization/58b0e1c2-3f6a-4c2e-9d3b-2f0a1b2c3d4e/logs/2c26b46b68ffc68f"
	// Note: The grammar is
	// `<type>:organization/<organization id>/<kind>/<selector>`.
	// For telemetry resources the type is `telemetryresource` and the selector
	// is replaced by its SHA-256 hash truncated to 16 hexadecimal characters.
	SigNozResourceObjectKey = attribute.Key("signoz.resource.object")

	// SigNozResourceTargetIDKey is the attribute Key conforming to the
	// "signoz.resource.target.id" semantic conventions. It represents the unique
	// identifier of the counterpart resource in a relation.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "7c1d9e40-8a2b-4f6d-b1c3-9e5f7a0d2b4c"
	// Note: Set on operations that relate two resources, for example attach and
	// detach. Follows the same rules as `signoz.resource.id`.
	SigNozResourceTargetIDKey = attribute.Key("signoz.resource.target.id")

	// SigNozResourceTargetKindKey is the attribute Key conforming to the
	// "signoz.resource.target.kind" semantic conventions. It represents the kind of
	// the counterpart resource in a relation.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "dashboard"
	// Note: Set on operations that relate two resources, for example attach and
	// detach: `attach dashboard D to alert rule R` records R as the target.
	// Relations emit one event per side; in each event the
	// `signoz.resource.*` attributes describe the side the event is recorded
	// from and the `signoz.resource.target.*` attributes the other side.
	//
	// Takes the same values as `signoz.resource.kind`.
	SigNozResourceTargetKindKey = attribute.Key("signoz.resource.target.kind")

	// SigNozResourceTargetObjectKey is the attribute Key conforming to the
	// "signoz.resource.target.object" semantic conventions. It represents the
	// object reference of the counterpart resource in a relation.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples:
	// "telemetryresource:organization/58b0e1c2-3f6a-4c2e-9d3b-2f0a1b2c3d4e/traces/9f86d081884c7d65"
	// Note: Set on operations that relate two resources, for example attach and
	// detach. Follows the same grammar as `signoz.resource.object`.
	SigNozResourceTargetObjectKey = attribute.Key("signoz.resource.target.object")
)

// SigNozResourceID returns an attribute KeyValue conforming to the
// "signoz.resource.id" semantic conventions. It represents the unique identifier
// of the resource.
func SigNozResourceID(val string) attribute.KeyValue {
	return SigNozResourceIDKey.String(val)
}

// SigNozResourceKind returns an attribute KeyValue conforming to the
// "signoz.resource.kind" semantic conventions. It represents the kind of the
// resource.
func SigNozResourceKind(val string) attribute.KeyValue {
	return SigNozResourceKindKey.String(val)
}

// SigNozResourceObject returns an attribute KeyValue conforming to the
// "signoz.resource.object" semantic conventions. It represents the object
// reference of the resource in the authorization model.
func SigNozResourceObject(val string) attribute.KeyValue {
	return SigNozResourceObjectKey.String(val)
}

// SigNozResourceTargetID returns an attribute KeyValue conforming to the
// "signoz.resource.target.id" semantic conventions. It represents the unique
// identifier of the counterpart resource in a relation.
func SigNozResourceTargetID(val string) attribute.KeyValue {
	return SigNozResourceTargetIDKey.String(val)
}

// SigNozResourceTargetKind returns an attribute KeyValue conforming to the
// "signoz.resource.target.kind" semantic conventions. It represents the kind of
// the counterpart resource in a relation.
func SigNozResourceTargetKind(val string) attribute.KeyValue {
	return SigNozResourceTargetKindKey.String(val)
}

// SigNozResourceTargetObject returns an attribute KeyValue conforming to the
// "signoz.resource.target.object" semantic conventions. It represents the object
// reference of the counterpart resource in a relation.
func SigNozResourceTargetObject(val string) attribute.KeyValue {
	return SigNozResourceTargetObjectKey.String(val)
}
