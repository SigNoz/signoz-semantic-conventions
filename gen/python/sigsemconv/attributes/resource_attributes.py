# Code generated from semantic convention specification. DO NOT EDIT.

from typing import Final

SIGNOZ_RESOURCE_ID: Final = "signoz.resource.id"
"""
Unique identifier of the resource.
Note: The id is unique within its kind and organization, not across kinds.
Its format depends on the kind; most platform objects use a UUID.

The value `*` denotes the collection of all resources of the kind and
is used when an operation applies to the collection rather than to a
single resource, for example a list operation.
"""

SIGNOZ_RESOURCE_KIND: Final = "signoz.resource.kind"
"""
Kind of the resource.
Note: Kind values are lowercase, with words joined by hyphens.

Known values include metaresource kinds such as `dashboard` and
`alert-rule`; telemetry resource kinds `logs`, `traces`, `metrics`,
`meter-metrics`, `logs-field` and `traces-field`; and platform object
kinds such as `ingestion-key`, `ingestion-limit`, `license`,
`subscription` and `deployment-host`. The set is not final.
"""

SIGNOZ_RESOURCE_OBJECT: Final = "signoz.resource.object"
"""
Object reference of the resource in the authorization model.
Note: The grammar is `<type>:organization/<organization id>/<kind>/<selector>`.
For telemetry resources the type is `telemetryresource` and the selector
is replaced by its SHA-256 hash truncated to 16 hexadecimal characters.
"""

SIGNOZ_RESOURCE_TARGET_ID: Final = "signoz.resource.target.id"
"""
Unique identifier of the counterpart resource in a relation.
Note: Set on operations that relate two resources, for example attach and
detach. Follows the same rules as `signoz.resource.id`.
"""

SIGNOZ_RESOURCE_TARGET_KIND: Final = "signoz.resource.target.kind"
"""
Kind of the counterpart resource in a relation.
Note: Set on operations that relate two resources, for example attach and
detach: `attach dashboard D to alert rule R` records R as the target.
Relations emit one event per side; in each event the
`signoz.resource.*` attributes describe the side the event is recorded
from and the `signoz.resource.target.*` attributes the other side.

Takes the same values as `signoz.resource.kind`.
"""

SIGNOZ_RESOURCE_TARGET_OBJECT: Final = "signoz.resource.target.object"
"""
Object reference of the counterpart resource in a relation.
Note: Set on operations that relate two resources, for example attach and
detach. Follows the same grammar as `signoz.resource.object`.
"""
