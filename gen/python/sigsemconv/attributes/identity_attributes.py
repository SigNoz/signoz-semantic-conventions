# Code generated from semantic convention specification. DO NOT EDIT.

from typing import Final

SIGNOZ_ORGANIZATION_ID: Final = "signoz.organization.id"
"""
Unique identifier of the organization.
Note: The value is an opaque string: a UUID today, but the format may
change.
"""

SIGNOZ_PRINCIPAL_AUTHN_METHOD: Final = "signoz.principal.authn.method"
"""
Method by which the principal identified itself on the request.
Note: The SigNoz platform defines the values, such as `tokenizer` (a user's
sign-in credential) and `api_key` (a service account API key).
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
