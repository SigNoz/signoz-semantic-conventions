// Code generated from semantic convention specification. DO NOT EDIT.

package sigsemconv

import "go.opentelemetry.io/otel/attribute"

// Area: audit
const (
	// SigNozAuditActionCategoryKey is the attribute Key conforming to the
	// "signoz.audit.action_category" semantic conventions. It represents the
	// category of the audited operation.
	//
	// Type: Enum
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "configuration_changes"
	// Note: Values are audit event categories of IEC 62443-3-3 SR 2.8, lowercased
	// with words joined by underscores.
	SigNozAuditActionCategoryKey = attribute.Key("signoz.audit.action_category")

	// SigNozAuditOutcomeKey is the attribute Key conforming to the
	// "signoz.audit.outcome" semantic conventions. It represents the outcome of the
	// audited operation.
	//
	// Type: Enum
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "success"
	// Note: The set is complete.
	//
	// An operation served over HTTP counts as a success when the response
	// status code is 2xx or 3xx. A failed operation also carries the upstream
	// `exception.*` attributes describing the failure.
	SigNozAuditOutcomeKey = attribute.Key("signoz.audit.outcome")

	// SigNozAuditVerbKey is the attribute Key conforming to the "signoz.audit.verb"
	// semantic conventions. It represents the operation the principal performed on
	// the resource.
	//
	// Type: string
	// RequirementLevel: Recommended
	// Stability: Development
	//
	// Examples: "update"
	// Note: Known values: `create`, `update`, `delete`, `attach`, `detach`.
	//
	// `attach` and `detach` relate two resources and are recorded once from
	// each side, with `signoz.resource.target.*` describing the other side.
	SigNozAuditVerbKey = attribute.Key("signoz.audit.verb")
)

// SigNozAuditVerb returns an attribute KeyValue conforming to the
// "signoz.audit.verb" semantic conventions. It represents the operation the
// principal performed on the resource.
func SigNozAuditVerb(val string) attribute.KeyValue {
	return SigNozAuditVerbKey.String(val)
}

// Enum values for signoz.audit.action_category
var (
	// A change to who may do what, such as a role assignment or a credential
	// change.
	//
	// Stability: development
	SigNozAuditActionCategoryAccessControl = SigNozAuditActionCategoryKey.String("access_control")
	// A change to a metaresource or a platform setting, such as a dashboard or an
	// alert rule change.
	//
	// Stability: development
	SigNozAuditActionCategoryConfigurationChanges = SigNozAuditActionCategoryKey.String("configuration_changes")
)

// Enum values for signoz.audit.outcome
var (
	// The operation succeeded.
	//
	// Stability: development
	SigNozAuditOutcomeSuccess = SigNozAuditOutcomeKey.String("success")
	// The operation failed.
	//
	// Stability: development
	SigNozAuditOutcomeFailure = SigNozAuditOutcomeKey.String("failure")
)

// Event names
const (
	// SigNozAuditAttachEventName is the name of the event conforming to the
	// "signoz.audit.attach" semantic conventions.
	//
	// Records the attachment of one SigNoz resource to another.
	//
	// Stability: Development
	SigNozAuditAttachEventName = "signoz.audit.attach"

	// SigNozAuditCreateEventName is the name of the event conforming to the
	// "signoz.audit.create" semantic conventions.
	//
	// Records the creation of a SigNoz resource.
	//
	// Stability: Development
	SigNozAuditCreateEventName = "signoz.audit.create"

	// SigNozAuditDeleteEventName is the name of the event conforming to the
	// "signoz.audit.delete" semantic conventions.
	//
	// Records the deletion of a SigNoz resource.
	//
	// Stability: Development
	SigNozAuditDeleteEventName = "signoz.audit.delete"

	// SigNozAuditDetachEventName is the name of the event conforming to the
	// "signoz.audit.detach" semantic conventions.
	//
	// Records the detachment of one SigNoz resource from another.
	//
	// Stability: Development
	SigNozAuditDetachEventName = "signoz.audit.detach"

	// SigNozAuditUpdateEventName is the name of the event conforming to the
	// "signoz.audit.update" semantic conventions.
	//
	// Records the modification of a SigNoz resource.
	//
	// Stability: Development
	SigNozAuditUpdateEventName = "signoz.audit.update"
)
