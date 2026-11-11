//----------------------------------------------------------------------------
// Code generated from semantic convention specification. DO NOT EDIT.
//----------------------------------------------------------------------------

'use strict';

Object.defineProperty(exports, '__esModule', { value: true });

/**
 * Unique identifier of the organization.
 *
 * @example 58b0e1c2-3f6a-4c2e-9d3b-2f0a1b2c3d4e
 *
 * @note The value is an opaque string: a UUID today, but the format may
 * change.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_ORGANIZATION_ID = 'signoz.organization.id';

/**
 * Method by which the request was identified.
 *
 * @example api_key
 *
 * @note Describes the request, not the principal: the same principal can use
 * different methods on different requests.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_PRINCIPAL_AUTHN_METHOD = 'signoz.principal.authn.method';

/**
 * Enum value "anonymous" for attribute {@link ATTR_SIGNOZ_PRINCIPAL_AUTHN_METHOD}.
 *
 * The request carried no credentials.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SIGNOZ_PRINCIPAL_AUTHN_METHOD_VALUE_ANONYMOUS = "anonymous";

/**
 * Enum value "api_key" for attribute {@link ATTR_SIGNOZ_PRINCIPAL_AUTHN_METHOD}.
 *
 * A service account API key, sent in the `SigNoz-Api-Key` header.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SIGNOZ_PRINCIPAL_AUTHN_METHOD_VALUE_API_KEY = "api_key";

/**
 * Enum value "impersonation" for attribute {@link ATTR_SIGNOZ_PRINCIPAL_AUTHN_METHOD}.
 *
 * Impersonation mode, in which every request acts as the root user and no credential is checked.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SIGNOZ_PRINCIPAL_AUTHN_METHOD_VALUE_IMPERSONATION = "impersonation";

/**
 * Enum value "internal" for attribute {@link ATTR_SIGNOZ_PRINCIPAL_AUTHN_METHOD}.
 *
 * A call made by SigNoz itself, such as an alert rule evaluation.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SIGNOZ_PRINCIPAL_AUTHN_METHOD_VALUE_INTERNAL = "internal";

/**
 * Enum value "tokenizer" for attribute {@link ATTR_SIGNOZ_PRINCIPAL_AUTHN_METHOD}.
 *
 * A bearer token issued to a user at sign-in.
 *
 * @experimental This enum value is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.SIGNOZ_PRINCIPAL_AUTHN_METHOD_VALUE_TOKENIZER = "tokenizer";

/**
 * Email address of the principal.
 *
 * @example jane.doe@example.com
 *
 * @note > [!WARNING]
 *
 * > This attribute may contain sensitive personally identifiable
 * > information (PII).
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_PRINCIPAL_EMAIL = 'signoz.principal.email';

/**
 * Unique identifier of the principal.
 *
 * @example 7c1d9e40-8a2b-4f6d-b1c3-9e5f7a0d2b4c
 *
 * @note Identifies a user or a service account, as indicated by
 * `signoz.principal.type`.
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_PRINCIPAL_ID = 'signoz.principal.id';

/**
 * Kind of the principal.
 *
 * @example user
 *
 * @note Known values: `user` (a human user of the SigNoz platform) and
 * `service_account` (a non-human principal that authenticates with an API
 * key).
 *
 * @experimental This attribute is experimental and is subject to breaking changes in minor releases of `sigsemconv`.
 */
exports.ATTR_SIGNOZ_PRINCIPAL_TYPE = 'signoz.principal.type';
