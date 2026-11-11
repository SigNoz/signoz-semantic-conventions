//----------------------------------------------------------------------------
// Code generated from semantic convention specification. DO NOT EDIT.
//----------------------------------------------------------------------------

'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

/**
 * Identifier of the key that authenticated the telemetry.
 *
 * @example 019d9a99-f8f0-7e50-8bb3-995599fa1f76
 *
 * @note Useful for identifying which key the telemetry was ingested under.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_WORKSPACE_KEY_ID = 'signoz.workspace.key.id';

/**
 * Identifier of the workspace under which the telemetry was ingested.
 *
 * @example acme-prod
 *
 * @note Useful for identifying which workspace the telemetry was ingested under.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_WORKSPACE_NAME = 'signoz.workspace.name';
