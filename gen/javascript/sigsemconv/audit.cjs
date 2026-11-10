//----------------------------------------------------------------------------
// Code generated from semantic convention specification. DO NOT EDIT.
//----------------------------------------------------------------------------

'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

/**
 * Category of the audited operation.
 *
 * @example configuration_changes
 *
 * @note Values are audit event categories of IEC 62443-3-3 SR 2.8, lowercased
 * with words joined by underscores.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_AUDIT_ACTION_CATEGORY = 'signoz.audit.action_category';

/**
 * Enum value "access_control" for attribute {@link ATTR_SIGNOZ_AUDIT_ACTION_CATEGORY}.
 *
 * A change to who may do what, such as a role assignment or a credential change.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SIGNOZ_AUDIT_ACTION_CATEGORY_VALUE_ACCESS_CONTROL = "access_control";

/**
 * Enum value "configuration_changes" for attribute {@link ATTR_SIGNOZ_AUDIT_ACTION_CATEGORY}.
 *
 * A change to a metaresource or a platform setting, such as a dashboard or an alert rule change.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SIGNOZ_AUDIT_ACTION_CATEGORY_VALUE_CONFIGURATION_CHANGES = "configuration_changes";

/**
 * Outcome of the audited operation.
 *
 * @example success
 *
 * @note The set is complete.
 *
 * An operation served over HTTP counts as a success when the response
 * status code is 2xx or 3xx. A failed operation also carries the upstream
 * `exception.*` attributes describing the failure.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_AUDIT_OUTCOME = 'signoz.audit.outcome';

/**
 * Enum value "failure" for attribute {@link ATTR_SIGNOZ_AUDIT_OUTCOME}.
 *
 * The operation failed.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SIGNOZ_AUDIT_OUTCOME_VALUE_FAILURE = "failure";

/**
 * Enum value "success" for attribute {@link ATTR_SIGNOZ_AUDIT_OUTCOME}.
 *
 * The operation succeeded.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SIGNOZ_AUDIT_OUTCOME_VALUE_SUCCESS = "success";

/**
 * Operation the principal performed on the resource.
 *
 * @example update
 *
 * @note Known values: `create`, `update`, `delete`, `attach`, `detach`.
 *
 * `attach` and `detach` relate two resources and are recorded once from
 * each side, with `signoz.resource.target.*` describing the other side.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_AUDIT_VERB = 'signoz.audit.verb';

/**
 * Records the attachment of one SigNoz resource to another.
 *
 * @note Recorded twice per attachment, once from each side of the relation,
 * after the operation has completed. For `attach dashboard D to alert rule
 * R`, one record describes D with `signoz.resource.*` and R with
 * `signoz.resource.target.*`, and the other record swaps them. The
 * timestamp is the time of the operation.
 *
 * Severity **SHOULD** be `INFO` (9) when `signoz.audit.outcome` is `success`
 * and `ERROR` (17) when it is `failure`.
 *
 * @experimental This event is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.EVENT_SIGNOZ_AUDIT_ATTACH = 'signoz.audit.attach';

/**
 * Records the creation of a SigNoz resource.
 *
 * @note Recorded once per creation, after the operation has completed. The
 * timestamp is the time of the operation.
 *
 * Severity **SHOULD** be `INFO` (9) when `signoz.audit.outcome` is `success`
 * and `ERROR` (17) when it is `failure`.
 *
 * @experimental This event is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.EVENT_SIGNOZ_AUDIT_CREATE = 'signoz.audit.create';

/**
 * Records the deletion of a SigNoz resource.
 *
 * @note Recorded once per deletion, after the operation has completed. The
 * timestamp is the time of the operation.
 *
 * Severity **SHOULD** be `INFO` (9) when `signoz.audit.outcome` is `success`
 * and `ERROR` (17) when it is `failure`.
 *
 * @experimental This event is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.EVENT_SIGNOZ_AUDIT_DELETE = 'signoz.audit.delete';

/**
 * Records the detachment of one SigNoz resource from another.
 *
 * @note Recorded twice per detachment, once from each side of the relation,
 * after the operation has completed. For `detach dashboard D from alert
 * rule R`, one record describes D with `signoz.resource.*` and R with
 * `signoz.resource.target.*`, and the other record swaps them. The
 * timestamp is the time of the operation.
 *
 * Severity **SHOULD** be `INFO` (9) when `signoz.audit.outcome` is `success`
 * and `ERROR` (17) when it is `failure`.
 *
 * @experimental This event is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.EVENT_SIGNOZ_AUDIT_DETACH = 'signoz.audit.detach';

/**
 * Records the modification of a SigNoz resource.
 *
 * @note Recorded once per modification, after the operation has completed. The
 * timestamp is the time of the operation.
 *
 * Severity **SHOULD** be `INFO` (9) when `signoz.audit.outcome` is `success`
 * and `ERROR` (17) when it is `failure`.
 *
 * @experimental This event is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.EVENT_SIGNOZ_AUDIT_UPDATE = 'signoz.audit.update';
