// Code generated from semantic convention specification. DO NOT EDIT.

package sigsemconv

import "go.opentelemetry.io/otel/attribute"

// Area: workspace
const (
	// SigNozWorkspaceKeyIDKey is the attribute Key conforming to the
	// "signoz.workspace.key.id" semantic conventions. It represents the identifier
	// of the key that authenticated the telemetry.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "019d9a99-f8f0-7e50-8bb3-995599fa1f76"
	// Note: Useful for identifying which key the telemetry was ingested under.
	SigNozWorkspaceKeyIDKey = attribute.Key("signoz.workspace.key.id")

	// SigNozWorkspaceNameKey is the attribute Key conforming to the
	// "signoz.workspace.name" semantic conventions. It represents the identifier of
	// the workspace under which the telemetry was ingested.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "acme-prod"
	// Note: Useful for identifying which workspace the telemetry was ingested
	// under.
	SigNozWorkspaceNameKey = attribute.Key("signoz.workspace.name")
)

// SigNozWorkspaceKeyID returns an attribute KeyValue conforming to the
// "signoz.workspace.key.id" semantic conventions. It represents the identifier
// of the key that authenticated the telemetry.
func SigNozWorkspaceKeyID(val string) attribute.KeyValue {
	return SigNozWorkspaceKeyIDKey.String(val)
}

// SigNozWorkspaceName returns an attribute KeyValue conforming to the
// "signoz.workspace.name" semantic conventions. It represents the identifier of
// the workspace under which the telemetry was ingested.
func SigNozWorkspaceName(val string) attribute.KeyValue {
	return SigNozWorkspaceNameKey.String(val)
}
