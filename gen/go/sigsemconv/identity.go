// Code generated from semantic convention specification. DO NOT EDIT.

package sigsemconv

import "go.opentelemetry.io/otel/attribute"

// Area: identity
const (
	// SigNozOrganizationIDKey is the attribute Key conforming to the
	// "signoz.organization.id" semantic conventions. It represents the unique
	// identifier of the organization.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "58b0e1c2-3f6a-4c2e-9d3b-2f0a1b2c3d4e"
	// Note: The value is an opaque string: a UUID today, but the format may
	// change.
	SigNozOrganizationIDKey = attribute.Key("signoz.organization.id")

	// SigNozPrincipalEmailKey is the attribute Key conforming to the
	// "signoz.principal.email" semantic conventions. It represents the email
	// address of the principal.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "jane.doe@example.com"
	// Note: > [!WARNING]
	//
	// > This attribute may contain sensitive personally identifiable
	// > information (PII).
	SigNozPrincipalEmailKey = attribute.Key("signoz.principal.email")

	// SigNozPrincipalIDKey is the attribute Key conforming to the
	// "signoz.principal.id" semantic conventions. It represents the unique
	// identifier of the principal.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "7c1d9e40-8a2b-4f6d-b1c3-9e5f7a0d2b4c"
	// Note: Identifies a user or a service account, as indicated by
	// `signoz.principal.type`.
	SigNozPrincipalIDKey = attribute.Key("signoz.principal.id")

	// SigNozPrincipalTypeKey is the attribute Key conforming to the
	// "signoz.principal.type" semantic conventions. It represents the kind of the
	// principal.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "user"
	// Note: Known values: `user` (a human user of the SigNoz platform) and
	// `service_account` (a non-human principal that authenticates with an API
	// key).
	SigNozPrincipalTypeKey = attribute.Key("signoz.principal.type")
)

// SigNozOrganizationID returns an attribute KeyValue conforming to the
// "signoz.organization.id" semantic conventions. It represents the unique
// identifier of the organization.
func SigNozOrganizationID(val string) attribute.KeyValue {
	return SigNozOrganizationIDKey.String(val)
}

// SigNozPrincipalEmail returns an attribute KeyValue conforming to the
// "signoz.principal.email" semantic conventions. It represents the email address
// of the principal.
func SigNozPrincipalEmail(val string) attribute.KeyValue {
	return SigNozPrincipalEmailKey.String(val)
}

// SigNozPrincipalID returns an attribute KeyValue conforming to the
// "signoz.principal.id" semantic conventions. It represents the unique
// identifier of the principal.
func SigNozPrincipalID(val string) attribute.KeyValue {
	return SigNozPrincipalIDKey.String(val)
}

// SigNozPrincipalType returns an attribute KeyValue conforming to the
// "signoz.principal.type" semantic conventions. It represents the kind of the
// principal.
func SigNozPrincipalType(val string) attribute.KeyValue {
	return SigNozPrincipalTypeKey.String(val)
}
