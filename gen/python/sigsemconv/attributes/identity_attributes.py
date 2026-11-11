# Code generated from semantic convention specification. DO NOT EDIT.

from enum import Enum
from typing import Final

SIGNOZ_ORGANIZATION_ID: Final = "signoz.organization.id"
"""
Unique identifier of the organization.
Note: The value is an opaque string: a UUID today, but the format may
change.
"""

SIGNOZ_PRINCIPAL_AUTHN_METHOD: Final = "signoz.principal.authn.method"
"""
Method by which the request was identified.
Note: Describes the request, not the principal: the same principal can use
different methods on different requests.
"""

SIGNOZ_PRINCIPAL_EMAIL: Final = "signoz.principal.email"
"""
Email address of the principal.
Note: > [!WARNING]
> This attribute may contain sensitive personally identifiable
> information (PII).
"""

SIGNOZ_PRINCIPAL_ID: Final = "signoz.principal.id"
"""
Unique identifier of the principal.
Note: Identifies a user or a service account, as indicated by
`signoz.principal.type`.
"""

SIGNOZ_PRINCIPAL_TYPE: Final = "signoz.principal.type"
"""
Kind of the principal.
Note: Known values: `user` (a human user of the SigNoz platform) and
`service_account` (a non-human principal that authenticates with an API
key).
"""


class SigNozPrincipalAuthnMethodValues(Enum):
    TOKENIZER = "tokenizer"
    """A bearer token issued to a user at sign-in."""
    API_KEY = "api_key"
    """A service account API key, sent in the `SigNoz-Api-Key` header."""
    ANONYMOUS = "anonymous"
    """The request carried no credentials."""
    INTERNAL = "internal"
    """A call made by SigNoz itself, such as an alert rule evaluation."""
    IMPERSONATION = "impersonation"
    """Impersonation mode, in which every request acts as the root user and no credential is checked."""
