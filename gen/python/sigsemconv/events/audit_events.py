# Code generated from semantic convention specification. DO NOT EDIT.

from typing import Final

SIGNOZ_AUDIT_ATTACH: Final = "signoz.audit.attach"
"""
Records the attachment of one SigNoz resource to another.
Note: Recorded twice per attachment, once from each side of the relation,
after the operation has completed. For `attach dashboard D to alert rule
R`, one record describes D with `signoz.resource.*` and R with
`signoz.resource.target.*`, and the other record swaps them. The
timestamp is the time of the operation.

Severity SHOULD be `INFO` (9) when `signoz.audit.outcome` is `success`
and `ERROR` (17) when it is `failure`.
"""

SIGNOZ_AUDIT_CREATE: Final = "signoz.audit.create"
"""
Records the creation of a SigNoz resource.
Note: Recorded once per creation, after the operation has completed. The
timestamp is the time of the operation.

Severity SHOULD be `INFO` (9) when `signoz.audit.outcome` is `success`
and `ERROR` (17) when it is `failure`.
"""

SIGNOZ_AUDIT_DELETE: Final = "signoz.audit.delete"
"""
Records the deletion of a SigNoz resource.
Note: Recorded once per deletion, after the operation has completed. The
timestamp is the time of the operation.

Severity SHOULD be `INFO` (9) when `signoz.audit.outcome` is `success`
and `ERROR` (17) when it is `failure`.
"""

SIGNOZ_AUDIT_DETACH: Final = "signoz.audit.detach"
"""
Records the detachment of one SigNoz resource from another.
Note: Recorded twice per detachment, once from each side of the relation,
after the operation has completed. For `detach dashboard D from alert
rule R`, one record describes D with `signoz.resource.*` and R with
`signoz.resource.target.*`, and the other record swaps them. The
timestamp is the time of the operation.

Severity SHOULD be `INFO` (9) when `signoz.audit.outcome` is `success`
and `ERROR` (17) when it is `failure`.
"""

SIGNOZ_AUDIT_UPDATE: Final = "signoz.audit.update"
"""
Records the modification of a SigNoz resource.
Note: Recorded once per modification, after the operation has completed. The
timestamp is the time of the operation.

Severity SHOULD be `INFO` (9) when `signoz.audit.outcome` is `success`
and `ERROR` (17) when it is `failure`.
"""
