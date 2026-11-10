//----------------------------------------------------------------------------
// Code generated from semantic convention specification. DO NOT EDIT.
//----------------------------------------------------------------------------

'use strict';

/**
 * The schema URL that matches the version of the semantic conventions that
 * this package defines.
 */
const SCHEMA_URL = 'https://signoz.io/schemas/otel/0.1.0-unreleased';

module.exports = {
  SCHEMA_URL,
  ...require('./audit.cjs'),
  ...require('./genai.cjs'),
  ...require('./identity.cjs'),
  ...require('./meter.cjs'),
  ...require('./resource.cjs'),
  ...require('./spanmetrics.cjs'),
  ...require('./workspace.cjs'),
};

Object.defineProperty(module.exports, '__esModule', { value: true });
