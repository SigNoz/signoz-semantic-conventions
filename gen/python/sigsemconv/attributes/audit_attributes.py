# Code generated from semantic convention specification. DO NOT EDIT.

from enum import Enum
from typing import Final

SIGNOZ_AUDIT_ACTION_CATEGORY: Final = "signoz.audit.action_category"
"""
Category of the audited operation.
Note: Values are audit event categories of IEC 62443-3-3 SR 2.8, lowercased
with words joined by underscores.
"""

SIGNOZ_AUDIT_OUTCOME: Final = "signoz.audit.outcome"
"""
Outcome of the audited operation.
Note: The set is complete.

An operation served over HTTP counts as a success when the response
status code is 2xx or 3xx. A failed operation also carries the upstream
`exception.*` attributes describing the failure.
"""

SIGNOZ_AUDIT_VERB: Final = "signoz.audit.verb"
"""
Operation the principal performed on the resource.
Note: Known values: `create`, `update`, `delete`, `attach`, `detach`.

`attach` and `detach` relate two resources and are recorded once from
each side, with `signoz.resource.target.*` describing the other side.
"""


class SignozAuditActionCategoryValues(Enum):
    ACCESS_CONTROL = "access_control"
    """A change to who may do what, such as a role assignment or a credential change."""
    CONFIGURATION_CHANGES = "configuration_changes"
    """A change to a metaresource or a platform setting, such as a dashboard or an alert rule change."""


class SignozAuditOutcomeValues(Enum):
    SUCCESS = "success"
    """The operation succeeded."""
    FAILURE = "failure"
    """The operation failed."""
