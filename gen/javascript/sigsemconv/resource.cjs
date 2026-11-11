//----------------------------------------------------------------------------
// Code generated from semantic convention specification. DO NOT EDIT.
//----------------------------------------------------------------------------

'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

/**
 * Unique identifier of the resource.
 *
 * @example 3f8a2c14-9b7d-4e1f-a6c3-5d8e2f0b1a4c
 * @example *
 *
 * @note The id is unique within its kind and organization, not across kinds.
 * Its format depends on the kind; most platform objects use a UUID.
 *
 * The value `*` denotes the collection of all resources of the kind and
 * is used when an operation applies to the collection rather than to a
 * single resource, for example a list operation.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_RESOURCE_ID = 'signoz.resource.id';

/**
 * Kind of the resource.
 *
 * @example ingestion-key
 *
 * @note Kind values are lowercase, with words joined by hyphens.
 *
 * Known values include metaresource kinds such as `dashboard` and
 * `alert-rule`; telemetry resource kinds `logs`, `traces`, `metrics`,
 * `meter-metrics`, `logs-field` and `traces-field`; and platform object
 * kinds such as `ingestion-key`, `ingestion-limit`, `license`,
 * `subscription` and `deployment-host`. The set is not final.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_RESOURCE_KIND = 'signoz.resource.kind';

/**
 * Object reference of the resource in the authorization model.
 *
 * @example telemetryresource:organization/58b0e1c2-3f6a-4c2e-9d3b-2f0a1b2c3d4e/logs/2c26b46b68ffc68f
 *
 * @note The grammar is `<type>:organization/<organization id>/<kind>/<selector>`.
 * For telemetry resources the type is `telemetryresource` and the selector
 * is replaced by its SHA-256 hash truncated to 16 hexadecimal characters.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_RESOURCE_OBJECT = 'signoz.resource.object';

/**
 * Unique identifier of the counterpart resource in a relation.
 *
 * @example 7c1d9e40-8a2b-4f6d-b1c3-9e5f7a0d2b4c
 *
 * @note Set on operations that relate two resources, for example attach and
 * detach. Follows the same rules as `signoz.resource.id`.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_RESOURCE_TARGET_ID = 'signoz.resource.target.id';

/**
 * Kind of the counterpart resource in a relation.
 *
 * @example dashboard
 *
 * @note Set on operations that relate two resources, for example attach and
 * detach: `attach dashboard D to alert rule R` records R as the target.
 * Relations emit one event per side; in each event the
 * `signoz.resource.*` attributes describe the side the event is recorded
 * from and the `signoz.resource.target.*` attributes the other side.
 *
 * Takes the same values as `signoz.resource.kind`.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_RESOURCE_TARGET_KIND = 'signoz.resource.target.kind';

/**
 * Object reference of the counterpart resource in a relation.
 *
 * @example telemetryresource:organization/58b0e1c2-3f6a-4c2e-9d3b-2f0a1b2c3d4e/traces/9f86d081884c7d65
 *
 * @note Set on operations that relate two resources, for example attach and
 * detach. Follows the same grammar as `signoz.resource.object`.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_RESOURCE_TARGET_OBJECT = 'signoz.resource.target.object';
